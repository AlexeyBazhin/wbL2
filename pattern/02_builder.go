package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/


/*
Определение
Порождающий паттерн, который позволяет поэтапно создавать составные объекты. 

Подробнее
Билдер позволяет создавать иммутабельные объекты на основе одного и того же процесса построения объекта. 
Решает проблему антипаттерна “телескопический конструктор”, т.к. не раздувает конструктор, а позволяет поэтапно наполнять объект с помощью методов.

Применяется когда
- Процесс **создания** нового объекта не должен зависеть от того, из каких частей этот объект состоит и как эти части связаны между собой
- Нужно получать различные вариации объекта в процессе создания - т.е. упростить создание различных представлений объекта.
*/

type (
	Bread struct {
		flour    string
		roasting string
		salt     int
		addition string
	}

	BreadBuilder interface {
		// NewBreadBuilder() BreadBuilder
		SetRoasting() BreadBuilder
		SetSalt() BreadBuilder
		SetAddition() BreadBuilder
		Build() *Bread
	}

	WheatBreadBuilder struct {
		Bread
	}
	RyeBreadBuilder struct {
		Bread
	}
	Baker struct {
		BreadBuilder
	}
)

func (bread *Bread) String() string {
	return fmt.Sprintf("Хлеб из %v, прожарка %v, соль: %v, добавки: %v", bread.flour, bread.roasting, bread.salt, bread.addition)
}

func NewWheatBreadBuilder() BreadBuilder {
	return &WheatBreadBuilder{
		Bread: Bread{
			flour: "wheat flour",
		},
	}
}

func (wheatBreadBuilder *WheatBreadBuilder) SetRoasting() BreadBuilder {
	wheatBreadBuilder.roasting = "medium"
	return wheatBreadBuilder
}

func (wheatBreadBuilder *WheatBreadBuilder) SetSalt() BreadBuilder {
	wheatBreadBuilder.salt = 20
	return wheatBreadBuilder
}

func (wheatBreadBuilder *WheatBreadBuilder) SetAddition() BreadBuilder {
	wheatBreadBuilder.addition = "provencal herbs"
	return wheatBreadBuilder
}

func (wheatBreadBuilder *WheatBreadBuilder) Build() *Bread {
	return &wheatBreadBuilder.Bread
}

func NewRyeBreadBuilder() BreadBuilder {
	return &RyeBreadBuilder{
		Bread: Bread{
			flour: "rye flour",
		},
	}
}

func (ryeBreadBuilder *RyeBreadBuilder) SetRoasting() BreadBuilder {
	ryeBreadBuilder.roasting = "well done"
	return ryeBreadBuilder
}

func (ryeBreadBuilder *RyeBreadBuilder) SetSalt() BreadBuilder {
	ryeBreadBuilder.salt = 12
	return ryeBreadBuilder
}

func (ryeBreadBuilder *RyeBreadBuilder) SetAddition() BreadBuilder {
	ryeBreadBuilder.addition = "oats"
	return ryeBreadBuilder
}

func (ryeBreadBuilder *RyeBreadBuilder) Build() *Bread {
	return &ryeBreadBuilder.Bread
}

// Director
func (baker *Baker) Bake() *Bread {
	return baker.BreadBuilder.
		SetRoasting().
		SetSalt().
		SetAddition().
		Build()
}

func main2() {
	baker := &Baker{
		BreadBuilder: NewWheatBreadBuilder(),
	}
	fmt.Println(baker.Bake())
	baker.BreadBuilder = NewRyeBreadBuilder()
	fmt.Println(baker.Bake())
}
