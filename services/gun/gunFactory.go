package gun

import "fmt"

func GetGun(gunType string) (iGun, error) {
    if gunType == "ak47" {
        return NewAk47(), nil
    }
    return nil, fmt.Errorf("Wrong gun type passed")
}