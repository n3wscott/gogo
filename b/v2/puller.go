package v2


type Puller struct {

}

func (p *Puller) Pull() string {
	return "b @ v2"
}