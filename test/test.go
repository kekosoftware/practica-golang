package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Define la estructura del JSON
type Data struct {
	MenuItems []MenuItem `json:"menu_items"`
}

// Define la estructura de un elemento del menú
type MenuItem struct {
	Label    string              `json:"label"`
	Children []MenuItem          `json:"children,omitempty"`
	Key      string              `json:"key,omitempty"`
}

func main() {
	// Leer el contenido del archivo JSON
	jsonData, err := ioutil.ReadFile("menu.json")
	if err != nil {
		fmt.Println("Error al leer el archivo JSON:", err)
		return
	}

	// Decodificar el JSON en la estructura de datos
	var data Data
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error al decodificar el JSON:", err)
		return
	}

	// Mostrar el menú
	selectedKeys := make([]string, 0)
	showMenu(data.MenuItems, selectedKeys)

	// Leer la clave correspondiente a las selecciones del usuario
	key := readKey(data.MenuItems, selectedKeys)
	fmt.Println("Clave seleccionada:", key)
}

// Mostrar el menú en la consola
func showMenu(menuItems []MenuItem, selectedKeys []string) {
	fmt.Println("Menú:")
	for i, item := range menuItems {
		fmt.Printf("%d. %s\n", i+1, item.Label)
	}
	fmt.Println("0. Seleccionar")

	// Mostrar selecciones anteriores
	fmt.Println("Selecciones anteriores:")
	for _, key := range selectedKeys {
		fmt.Printf("- %s\n", key)
	}
}

// Leer la selección del usuario y devolver la clave correspondiente
func readKey(menuItems []MenuItem, selectedKeys []string) string {
	var choice int
	for {
		fmt.Print("Ingrese su elección: ")
		fmt.Scanln(&choice)

		if choice == 0 {
			if len(selectedKeys) > 0 {
				selectedKeys = selectedKeys[:len(selectedKeys)-1]
			}
		} else if choice >= 1 && choice <= len(menuItems) {
			selectedKeys = append(selectedKeys, menuItems[choice-1].Label)
			if len(menuItems[choice-1].Children) > 0 {
				showMenu(menuItems[choice-1].Children, selectedKeys)
			} else {
				return menuItems[choice-1].Key
			}
		} else {
			fmt.Println("Selección inválida. Por favor, elija una opción válida.")
		}
	}
}