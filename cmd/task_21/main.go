package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

/*
Suppose there is a third-party library that we don't want to edit
*/
type FSNodeKind int

const (
	File FSNodeKind = iota
	Dir
)

type FSNode struct {
	Name     string
	Kind     FSNodeKind
	Children []FSNode
	Content  []byte
}

type FHS interface {
	Get(string) (FSNode, error)
}

type FHSSystem struct {
	BinDir map[string]FSNode
}

// Get returns the executable by its name
func (fhs *FHSSystem) Get(name string) (FSNode, error) {
	node, found := fhs.BinDir[name]
	if !found {
		return FSNode{}, fmt.Errorf("/bin/%s not found", name)
	}
	return node, nil
}

// some programs make assumptions on the structure of the file system -
// they expect it to comply to the Filesystem Hierarchy Standard
type FHSProgram struct {
	// names of programs that FHSProgram depends on
	Dependencies []string
}

// Run finds dependencies and checks them
func (p *FHSProgram) Run(system FHS) error {
	for _, dependency := range p.Dependencies {
		node, err := system.Get(dependency)
		if err != nil {
			return err
		}
		switch node.Kind {
		case Dir:
			return fmt.Errorf("%s is a directory, not a file", dependency)
		case File:
			continue
		}
	}
	return nil
}

type User string

// NixOS store only a `sh` symlink in the bin dir
type NixOSFileSystem struct {
	UserPrograms map[User](map[string]FSNode)
}

// adapter implements FHS interface
type NixOSFHSAdapter struct {
	user User
	host *NixOSFileSystem
}

func (a *NixOSFHSAdapter) Get(name string) (FSNode, error) {
	node, found := a.host.UserPrograms[a.user][name]
	if !found {
		err := fmt.Errorf("/etc/profiles/per-user/%s/bin/%s not found", a.user, name)
		return FSNode{}, err
	}
	return node, nil
}

func (nixos *NixOSFileSystem) WrapFHS(user User) (NixOSFHSAdapter, error) {
	_, found := nixos.UserPrograms[user]
	if !found {
		return NixOSFHSAdapter{}, fmt.Errorf("user %s does not exist", user)
	}
	return NixOSFHSAdapter{user, nixos}, nil
}

func main() {
	typicalUnixFS := &FHSSystem{
		BinDir: map[string]FSNode{
			"program": {
				Name:    "program",
				Kind:    File,
				Content: []byte("program_machine_code"),
			},
		},
	}
	dependsOnProgram := FHSProgram{[]string{"program"}}
	// runs just fine on a typical unix file system
	if err := dependsOnProgram.Run(typicalUnixFS); err != nil {
		err = fmt.Errorf("error running dependsOnProgram: %w", err)
		fmt.Println(err)
	} else {
		fmt.Println("dependsOnProgram run successfully")
	}

	nixosFS := &NixOSFileSystem{
		UserPrograms: map[User]map[string]FSNode{
			"root": {
				"bash": {
					Name:     "bash",
					Kind:     File,
					Children: nil,
					Content:  []byte("bash_code"),
				},
			},
			"hofsiedge": {
				"program": {
					Name:     "program",
					Kind:     File,
					Children: nil,
					Content:  []byte("program_machine_code"),
				},
			},
		},
	}

	// wrapping NixOSFileSystem with an adapter
	adapter, err := nixosFS.WrapFHS("hofsiedge")
	if err != nil {
		err = fmt.Errorf("could not FHS wrap a NixOS file system: %w", err)
		fmt.Println(err)
		return
	}
	dependsOnProgram.Run(&adapter)
	if err := dependsOnProgram.Run(typicalUnixFS); err != nil {
		err = fmt.Errorf("error running dependsOnProgram: %w", err)
		fmt.Println(err)
	} else {
		fmt.Println("dependsOnProgram run successfully")
	}
}
