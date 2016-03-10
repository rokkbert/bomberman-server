package main

import "bomberman-server/helper"

type GameMap struct {
	game   *Game
	size   int
	fields [][]*Field
}

func NewGameMap(size int) *GameMap {
	f := createFields(size)

	return &GameMap{
		size:   size,
		fields: f,
	}
}

func (gameMap *GameMap) field(vCode int, hCode int) *Field {
	return gameMap.fields[vCode][hCode]
}

func createFields(size int) [][]*Field {
	//horizontalFieldCodes := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	//verticalFieldCodes := []string{"01","02","03","04","05","06","07","08","09","10","11","12","13","14","15","16","17","18","19","20","21","22","23","24","25","26"}

	fields := make([][]*Field, size)
	for i := range fields {
		fields[i] = make([]*Field, size)
		for j := range fields[i] {
			fields[i][j] = NewField(i, j)
		}
	}

	// place walls and specials on the game map
	walls := 20
	specials := 5

	// place walls
	for i := 0; i <= walls; i++ {
		randomHorzitontalFieldCode := helper.RandomNumber(0, size)
		randomVerticalFieldCode := helper.RandomNumber(0, size)

		// TODO: prüfen ob auf dem Feld schon so ein Element liegt

		fields[randomVerticalFieldCode][randomHorzitontalFieldCode].setWall(true)
	}

	// place specials
	for i := 0; i <= specials; i++ {
		randomHorzitontalFieldCode := helper.RandomNumber(0, size)
		randomVerticalFieldCode := helper.RandomNumber(0, size)

		// TODO: prüfen ob auf dem Feld schon so ein Element liegt

		fields[randomVerticalFieldCode][randomHorzitontalFieldCode].setSpecial(true)
	}

	return fields
}

func (gm *GameMap) toString() string {
	mapString := "\n"
	// fmt.Println()
	for i := range gm.fields {
		for j := range gm.fields[i] {
			f := gm.fields[i][j]
			// h := f.horizontalFieldCode
			// v := f.verticalFieldCode

			if len(f.players) > 0 {
				// fmt.Printf("P")
				mapString += "P"
			} else if f.containsWall == true {
				// fmt.Printf("W")
				mapString += "W"
			} else if f.containsSpecial == true {
				// fmt.Printf("S")
				mapString += "S"
			} else {
				//fmt.Printf("_") //fmt.Printf("i %d, j %d", h, v) //fmt.Print(h + v)
				mapString += "_"
			}

			// fmt.Print("|")
			mapString += "|"
		}
		// fmt.Println()
		mapString += "\n"
	}

	//gm.game.mainChannel <- mapString
	return mapString
}
