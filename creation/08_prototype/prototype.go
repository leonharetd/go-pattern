package prototype

type Cloneable interface {
	Clone() CloneLab
}

type CloneLab struct {
    Lab map[string]string
}

func (l *CloneLab) set(key, value string) {
	l.Lab[key] = value
}

func (l *CloneLab) Clone() CloneLab {
	newMap := make(map[string]string, 0)
	for k,v := range l.Lab {
		newMap[k] = v
	  }
	return CloneLab{
		Lab: newMap,
	}
	
}