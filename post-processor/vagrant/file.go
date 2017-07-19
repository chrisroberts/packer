package vagrant

import (
	"fmt"
	"github.com/hashicorp/packer/packer"
	// "path/filepath"
)

type FileProvider struct{}

func (p *FileProvider) KeepInputArtifact() bool {
	return true
}

func (p *FileProvider) Process(ui packer.Ui, artifact packer.Artifact, dir string) (vagrantfile string, metadata map[string]interface{}, err error) {
	// Copy all of the original contents into the temporary directory
	for _, path := range artifact.Files() {
		ui.Message(fmt.Sprintf("Unpacking: %s", path))
		err = BoxToDir(dir, path, ui)
		if err != nil {
			return
		}
	}

	// if exist
	// vFile = os.Open(filepath.Join(dir, "Vagrantfile"))
	// defer vFile.Close()

	// vagrantfile = ioutil.ReadAll(vFile)

	// // if exist
	// mFile = os.Open(filepath.join(dir, "metadata.json"))
	// defer mFile.Close()

	// Create the metadata
	metadata = map[string]interface{}{"provider": "file"}

	return
}
