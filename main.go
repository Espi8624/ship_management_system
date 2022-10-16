package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ShipInfo struct {
	ShipType     string
	Country      string
	ShipName     string
	Manufacturer string
	TotalTonnage int
}

func dataAppender(list *[]ShipInfo, shipe_type string, country string, name string, manuf string, tonnage int) {
	*list = append((*list), ShipInfo{
		strings.TrimSpace(shipe_type),
		strings.TrimSpace(country),
		strings.TrimSpace(name),
		strings.TrimSpace(manuf),
		tonnage,
	})
}

func dataScannInt() int {
	stdin := bufio.NewReader(os.Stdin)
	var input int

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	}

	return input
}

func dataScanStr() string {
	stdin := bufio.NewReader(os.Stdin)

	input, err := stdin.ReadString('\n')
	if err != nil {
		fmt.Println("error message: ", err)
	}

	return input
}

var page int
var command int

func main() {
	shipData := []ShipInfo{}

	for {
		switch page {
		case 0:
			fmt.Print(`
					- 선박 등록 프로그램 -
			선박 등록 프로그램입니다. 선박을 등록, 조회 할수 있습니다.
	
			1. 조회
			2. 등록
			3. 삭제
			4. 종료
	
			입력: `)

			command = dataScannInt()
			if command == 1 {
				page = 1
			} else if command == 2 {
				page = 2
			} else if command == 3 {
				page = 3
			} else if command == 4 {
				page = 4
			} else {
				fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
			}

		case 1:
			fmt.Println("조회 화면입니다. 메인 화면으로 돌아가려면 '0 + enter' 를 입력해주세요.")
			if len(shipData) == 0 {
				fmt.Println("등록된 선박이 없습니다.")
			} else {
				for i := 0; i < len(shipData); i++ {
					fmt.Printf("No: %d => 선박타입: %v, 선박국적: %v, 선박명: %v, 조선사: %v, 선박톤수: %v\n",
						i, shipData[i].ShipType, shipData[i].Country, shipData[i].ShipName, shipData[i].Manufacturer, shipData[i].TotalTonnage)
				}
			}

			command = dataScannInt()
			if command == 0 {
				page = 0
			}

		case 2:
			fmt.Println("등록 화면입니다. 메인 화면으로 돌아가려면 '0 + enter', 등록을 시작하려면 '1 + enter' 을 눌러주세요.")
			command = dataScannInt()

			if command == 0 {
				page = 0
			} else if command == 1 {
				var shipType, country, shipName, manufacturer string
				var totalTonnage int

				fmt.Print("Please input ship type code: ")
				shipType = dataScanStr()

				fmt.Print("Please input country code: ")
				country = dataScanStr()

				fmt.Print("Please input ship name: ")
				shipName = dataScanStr()

				fmt.Print("Please input manufacturer name: ")
				manufacturer = dataScanStr()

				fmt.Print("Please input total tonnage of ship: ")
				totalTonnage = dataScannInt()

				dataAppender(&shipData, shipType, country, shipName, manufacturer, totalTonnage)
			} else {
				fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
			}

		case 3:
			fmt.Println("삭제 화면입니다. 메인 화면으로 돌아가려면 '0 + enter' 를 입력해주세요. 삭제를 시작하려면 '1 + enter' 을 눌러주세요.")
			if len(shipData) == 0 {
				fmt.Println("등록된 선박이 없습니다.")
			} else {
				for i := 0; i < len(shipData); i++ {
					fmt.Printf("No: %d => 선박타입: %v, 선박국적: %v, 선박명: %v, 조선사: %v, 선박톤수: %v\n",
						i+1, shipData[i].ShipType, shipData[i].Country, shipData[i].ShipName, shipData[i].Manufacturer, shipData[i].TotalTonnage)
				}
			}

			command = dataScannInt()
			if command == 0 {
				page = 0
			} else if command == 1 {
				fmt.Println("삭제할 선박의 번호를 입력해주세요.")
				command = dataScannInt()
				deleteData := append(shipData[:(command)], shipData[(command)+1:]...)
				shipData = deleteData
			}

		default:
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
		}

		if page == 4 {
			fmt.Println("프로그램을 종료합니다.")
			break
		}
	}
}
