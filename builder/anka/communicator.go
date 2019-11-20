package anka

import (
	"fmt"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/veertuinc/packer-builder-veertu-anka/client"
	"github.com/hashicorp/packer/packer"
)

type Communicator struct {
	Config  *Config
	Client  *client.Client
	HostDir string
	VMDir   string
	VMName  string
}

func (c *Communicator) Start(remote *packer.RemoteCmd) error {
	log.Printf("Communicator Start: %s", remote.Command)

	runner := client.NewRunner(client.RunParams{
		VMName:  c.VMName,
		Command: []string{remote.Command},
		Volume:  "",
		Stdout:  remote.Stdout,
		Stderr:  remote.Stderr,
		Stdin:   remote.Stdin,
	})

	if err := runner.Start(); err != nil {
		return err
	}

	go func() {
		err, exitCode := runner.Wait()
		if err != nil {
			log.Printf("Runner exited with error: %v", err)
		}
		remote.SetExited(exitCode)
	}()

	return nil

}

func (c *Communicator) Upload(dst string, src io.Reader, fi *os.FileInfo) error {
	log.Printf("Communicator Upload")
	// Create a temporary file to store the upload
	tempfile, err := ioutil.TempFile(c.HostDir, "upload")
	if err != nil {
		return err
	}
	defer os.Remove(tempfile.Name())

	// Copy the contents to the temporary file
	w, err := io.Copy(tempfile, src)
	if err != nil {
		return err
	}

	if fi != nil {
		tempfile.Chmod((*fi).Mode())
	}
	tempfile.Close()

	log.Printf("Created temp dir in %s", tempfile.Name())
	log.Printf("Copying %d bytes from %s to %s", w, tempfile.Name(), dst)

	log.Printf("Destination os %v", dst)

	err, _ = c.Client.Run(client.RunParams{
		VMName:  c.VMName,
		Command: []string{"cp", path.Base(tempfile.Name()), dst},
		Volume:  c.HostDir,
	})

	return err
}

func (c *Communicator) UploadDir(dst string, src string, exclude []string) error {
	// Create the temporary directory that will store the contents of "src"
	// for copying into the container.
	td, err := ioutil.TempDir(c.HostDir, "dirupload")
	if err != nil {
		return err
	}
	defer os.RemoveAll(td)

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relpath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		hostpath := filepath.Join(td, relpath)

		// If it is a directory, just create it
		if info.IsDir() {
			return os.MkdirAll(hostpath, info.Mode())
		}

		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			dest, err := os.Readlink(path)

			if err != nil {
				return err
			}

			return os.Symlink(dest, hostpath)
		}

		// It is a file, copy it over, including mode.
		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(hostpath)
		if err != nil {
			return err
		}
		defer dst.Close()

		log.Printf("Copying %s to %s", src.Name(), dst.Name())
		if _, err := io.Copy(dst, src); err != nil {
			return err
		}

		si, err := src.Stat()
		if err != nil {
			return err
		}

		return dst.Chmod(si.Mode())
	}

	// Copy the entire directory tree to the temporary directory
	if err := filepath.Walk(src, walkFn); err != nil {
		return err
	}

	// Determine the destination directory
	// containerSrc := filepath.Join(c.ContainerDir, filepath.Base(td))
	containerDst := dst
	if src[len(src)-1] != '/' {
		containerDst = filepath.Join(dst, filepath.Base(src))
	}

	log.Printf("from %#v to %#v", td, containerDst)

	// Make the directory, then copy into it
	command := fmt.Sprintf("set -e; mkdir -p %s; command cp -R %s/* %s",
		containerDst, filepath.Base(td), containerDst,
	)
	err, _ = c.Client.Run(client.RunParams{
		VMName:  c.VMName,
		Command: []string{"bash", "-c", command},
		Volume:  c.HostDir,
	})

	return err
}

func (c *Communicator) Download(src string, dst io.Writer) error {
	log.Printf("Downloading file from VM: %s", src)

	// Create a temporary file to store the download
	tempfile, err := ioutil.TempFile(c.HostDir, "download")
	if err != nil {
		return err
	}
	defer os.Remove(tempfile.Name())

	// Copy it to a local file mounted on shared fs
	err, _ = c.Client.Run(client.RunParams{
		VMName:  c.VMName,
		Command: []string{"cp", src, "./" + path.Base(tempfile.Name())},
		Volume:  c.HostDir,
	})

	log.Printf("Copying from %s to writer", tempfile.Name())
	w, err := io.Copy(dst, tempfile)
	if err != nil {
		return err
	}

	log.Printf("Copied %d bytes", w)
	return nil
}

func (c *Communicator) DownloadDir(src string, dst string, exclude []string) error {
	return errors.New("communicator.DownloadDir isn't implemented")
}
