package person

//대문자 public, 소문자 private
type Person struct {
	//접두문자를 대문자로 바꾸면  public이 됨
	name string
	age  int
}

//make pointer
func (instance *Person) SetDetails(name string, age int) {
	instance.name = name
	instance.age = age

}

func (p Person) GetName() string {
	return p.name
}
