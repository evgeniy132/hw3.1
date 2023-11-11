package main

import (
	"fmt"
)

type Scene struct {
	LeftLocation  *Scene
	RightLocation *Scene
	Description   string
	Win           bool
}

func (s *Scene) getNextLocation(walkLeft bool) *Scene {
	if walkLeft {
		return s.LeftLocation
	}
	return s.RightLocation
}

func NewScene(d string, Left *Scene, Right *Scene, Win bool) Scene {
	return Scene{LeftLocation: Left, RightLocation: Right, Description: d, Win: Win}
}

func (s *Scene) MoveNext(walkLeft bool) (*Scene, error) {
	nextScene := s.getNextLocation(walkLeft)
	if nextScene == nil {
		return nil,
			fmt.Errorf("hero dead")
	}

	return nextScene, nil
}

var (
	s_1 = NewScene("First location: Stephen opened his eyes near the cave and in front of him there were two paths: to the left into the path, and to the right the entrance to the cave", &s_2, &s_3, false) // айди передать как по листу + добавить описание. вибрать ласт сцену где он вин добавить еще 1 парметр вин или не вин
	s_2 = NewScene("Second Location: Stephen walked along a path near the forest and saw an animal", &s_6, &s_7, false)                                                                                       // дописать дискрипшн
	s_3 = NewScene("Third Location: Stephen went into a cave where it was cold and dark ", &s_7, nil, false)                                                                                                  //
	s_4 = NewScene("Fourth Location : Stephen found his way home, the hero won ", nil, nil, true)                                                                                                             //win локация
	s_5 = NewScene("Fifth Location: Stephen walked further and saw a safe on the left, and a path further on the right", nil, &s_4, false)                                                                    // поменять
	s_6 = NewScene("Sixth Location: Stephen found himself in a deserted camp and he has a choice: on the left is a bed on which he can lie down and rest, and on the right is a path to move on.", &s_5, nil, false)
	s_7 = NewScene("Seven Location: The hero fell into a trap. Choose which hand he can help himself with: with his left hand he simply tries to break the trap, and with his right hand he tries to get a knife and cut the trap", nil, &s_8, false)
	s_8 = NewScene("Eight Location: Stephen broke free and got on the right road", nil, &s_4, false)
)

func main() {

	Current := &s_1
	//fmt.Println("Stiven rodilsia")
	for {
		fmt.Println(Current.Description, "choose 1 or 2 , left = 1 , right = 2")
		var UserInput int
		fmt.Scan(&UserInput)
		fmt.Println("User input : ", UserInput)
		walkLeft := UserInput == 1

		Next, err := Current.MoveNext(walkLeft)
		if err != nil {
			fmt.Println("Next is nill = finish", err.Error())
			break
		}
		//тут нужна проверка на Win локацию если вин то написать
		if Next.Win {
			fmt.Println(Next.Description)
			break
		}

		Current = Next
	}

}
