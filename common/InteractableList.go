package common

type InteractableList struct {
	Items []Interactable
}

func (dq *InteractableList) Add(i Interactable) {
	dq.Items = append(dq.Items, i)
}

func (dq *InteractableList) Clear() {
	dq.Items = dq.Items[:0]
}
