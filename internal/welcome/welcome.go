package welcome

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

var (
	imageBG *ebiten.Image
)

func init() {
	var err error

	if imageBG, err = newImageFromFile("files/welcome_bg.png"); err != nil {
		log.Fatal(err)
	}
}

type Registration struct {
	ScreenWidth  int
	ScreenHeight int
	ui           *ebitenui.UI
}

func Show(screenWidth, screenHeight int, title string) {
	//face, _ := loadFont("files/fonts/NotoSans-Regular.ttf", 80)
	face, _ := loadFontTTF(80)

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(20)),
			widget.RowLayoutOpts.Spacing(20),
		)),
	)

	label1 := widget.NewText(
		widget.TextOpts.Text("Label 1 (NewText)", face, color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)

	// Add the first Text as a child of the container
	rootContainer.AddChild(label1)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle(title)

	ws := Registration{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		ui: &ebitenui.UI{
			Container: rootContainer,
		},
	}

	if err := ebiten.RunGame(&ws); err != nil {
		log.Fatal(err)
	}
}

func (r *Registration) Update() error {
	return nil
}

func (r *Registration) Draw(screen *ebiten.Image) {
	screen.DrawImage(imageBG, nil)
}

func (r *Registration) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return r.ScreenWidth, r.ScreenHeight
}
