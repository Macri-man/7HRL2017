

type teleport struct{

}

type Room struct{
  id int
  xpos,ypos int
  w,h int
  //items item[]
  //creatures Creatures[]

}

func makeRoom(){

}



type Dungeon struct{
  id string
  dungeon Room[]
  //items Item[]
  //creatures Creatures[]
}


func makeDungeon(numRooms int) Dungeon {
  r := rand.New(time.Now().UnixNano())
    for i := 0; i < numRooms; i++ {
      room := makeRoom(r.Intn(10))
    }
}
