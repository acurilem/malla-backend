package services

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/models"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const (
	fontSizeTitulo          = 8
	fontSizeEncabezadoTabla = 6
	fontSize                = 6
	fontSizeSimbologiaTexto = 6
	pageNumberLabel         = "Página {current} de {total}"
)

var (
	encabezadoTablaCentrado = props.Text{
		Size:  fontSizeEncabezadoTabla,
		Align: align.Center,
		Style: fontstyle.Bold,
		Top:   2.0,
	}

	encabezadoTabla = props.Text{
		Size:            fontSizeEncabezadoTabla,
		Style:           fontstyle.Bold,
		Left:            2.0,
		Top:             2.0,
		Right:           1.0,
		VerticalPadding: 1.0,
	}

	simbologiaTexto = props.Text{
		Size:            fontSizeSimbologiaTexto,
		Style:           fontstyle.Bold,
		Left:            2.0,
		Top:             2.0,
		Right:           1.0,
		VerticalPadding: 1.0,
	}

	tituloDocumento = props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Align: align.Center,
		Size:  fontSizeTitulo,
	}

	tituloTextoSemestre = props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Align: align.Center,
		Size:  7,
	}

	estiloSimbologia = props.Cell{
		BackgroundColor: &props.Color{
			Red:   243,
			Green: 243,
			Blue:  243},
		BorderColor: &props.Color{
			Red:   243,
			Green: 243,
			Blue:  243},
		//BorderType:      border.Full,
		BorderThickness: 0,
		LineStyle:       "solid",
	}

	estiloSimbologiaCelda = props.Cell{
		BackgroundColor: &props.Color{Red: 243, Green: 243, Blue: 243},
		BorderThickness: 0,
	}
	estiloSimbologiaCelda2 = props.Cell{
		BackgroundColor: &props.Color{Red: 230, Green: 230, Blue: 230},
		BorderThickness: 0,
	}
)

func GetPdfService(idMalla int) core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber(pageNumberLabel, props.SouthEast).
		WithMargins(10, 15, 10).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	if err := m.RegisterHeader(getPageHeader()); err != nil {
		log.Fatal(err.Error())
	}

	planes, err := GetPlanService(idMalla)
	if err != nil {
		return nil
	}
	fmt.Printf("Planes obtenidos: %v\n", planes)

	var idPlan uint
	for _, plan := range planes {

		// Convierte plan.CodTipoPlan a una cadena (string) para su uso
		idPlan = plan.CodTipoPlan

		titulo := fmt.Sprintf("%s - %s - %s - %s - %d - %s - %s",
			plan.CarreraTotal.Carrera.NomCarrera,
			plan.CarreraTotal.CodCarrera,
			plan.CarreraTotal.Especialidad.NomEspecialidad,
			plan.CarreraTotal.Mencion.NomMencion,
			plan.AnoPlan,
			plan.Regimen.NomRegimen,
			plan.TipoPlan.NomTipoPlan)

		m.AddRows(text.NewRow(10, titulo, tituloDocumento),
			row.New(5),
		)
	}

	cursos, err := GetMallaSinCursosComplementariosService(idMalla)
	if err != nil {
		return nil
	}

	sort.Slice(cursos, func(i, j int) bool {
		if cursos[i].NroAno == cursos[j].NroAno {
			return cursos[i].NroSem < cursos[j].NroSem
		}
		return cursos[i].NroAno < cursos[j].NroAno
	})

	grupos := make(map[int][]models.MallaPionero)
	for _, malla := range cursos {
		key := int(malla.NroAno)*10 + int(malla.NroSem)
		grupos[key] = append(grupos[key], malla)
	}

	var keys []int
	for key := range grupos {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	gray := getGrayColor()

	m.AddRow(7,
		text.NewCol(12, "SIMBOLOGIA", encabezadoTablaCentrado),
	).WithStyle(&props.Cell{BackgroundColor: gray})

	// si es TEL
	if idPlan == 4 {
		m.AddRow(7,
			text.NewCol(3, "DUR.", simbologiaTexto).WithStyle(&estiloSimbologiaCelda),
			text.NewCol(3, "DURACIÓN", simbologiaTexto).WithStyle(&estiloSimbologiaCelda2),
			text.NewCol(3, "*", simbologiaTexto).WithStyle(&estiloSimbologiaCelda),
			text.NewCol(3, "INSTITUCIONAL", simbologiaTexto).WithStyle(&estiloSimbologiaCelda2),
		)
	} else {
		m.AddRow(7,
			text.NewCol(1, "DUR.", simbologiaTexto).WithStyle(&estiloSimbologiaCelda),
			text.NewCol(2, "DURACIÓN", simbologiaTexto).WithStyle(&estiloSimbologiaCelda2),
			text.NewCol(1, "*", simbologiaTexto).WithStyle(&estiloSimbologiaCelda),
			text.NewCol(2, "INSTITUCIONAL", simbologiaTexto).WithStyle(&estiloSimbologiaCelda2),
			text.NewCol(1, "H.P", simbologiaTexto).WithStyle(&estiloSimbologiaCelda),
			text.NewCol(2, "HORAS PRESENCIALES", simbologiaTexto).WithStyle(&estiloSimbologiaCelda2),
			text.NewCol(1, "H.N.P.", simbologiaTexto).WithStyle(&estiloSimbologiaCelda),
			text.NewCol(2, "HORAS NO PRESENCIALES", simbologiaTexto).WithStyle(&estiloSimbologiaCelda2),
		)
	}

	m.AddRow(12,
		col.New(0),
		text.NewCol(12, "", encabezadoTabla),
	)

	//print idPlan
	fmt.Printf("El valor de idPlan es: %d\n", idPlan)

	for _, key := range keys {
		grupo := grupos[key]
		ano := key / 10
		semestre := key % 10

		var tituloGrupo string
		if ano == 6 && idMalla == 608 {
			tituloGrupo = fmt.Sprintf("Año 6 y Año 7")
		} else {
			tituloGrupo = fmt.Sprintf("Año: %d Semestre: %d", ano, semestre)
		}

		//Ordena alfabéticamente
		sort.Slice(grupo, func(i, j int) bool {
			return grupo[i].CursoPionero.CompetenciaAsignatura.NomCompAsig < grupo[j].CursoPionero.CompetenciaAsignatura.NomCompAsig
		})

		//add empty row except first iteration
		if key != keys[0] {
			m.AddRow(5)
		}

		//EMPRESARIAL
		if idPlan == 2 {
			m.AddRow(8,
				text.NewCol(2, "ASIGNATURA", encabezadoTablaCentrado),
				text.NewCol(2, "COMPETENCIA", encabezadoTablaCentrado),
				text.NewCol(1, "CT", encabezadoTablaCentrado),
				text.NewCol(1, "H.P.", encabezadoTablaCentrado),
				text.NewCol(1, "H.N.P.", encabezadoTablaCentrado),
				text.NewCol(1, "TIPO", encabezadoTablaCentrado),
				text.NewCol(1, "DUR.", encabezadoTablaCentrado),
				text.NewCol(1, "CURSO BÁSICO", encabezadoTablaCentrado),
				text.NewCol(2, "REQUISITOS", encabezadoTabla),
			).WithStyle(&props.Cell{BackgroundColor: gray})
		} else if idPlan == 4 {
			m.AddRow(8,
				text.NewCol(3, "ASIGNATURA", encabezadoTablaCentrado),
				text.NewCol(1, "T", encabezadoTablaCentrado),
				text.NewCol(1, "E", encabezadoTablaCentrado),
				text.NewCol(1, "L", encabezadoTablaCentrado),
				text.NewCol(1, "TIPO", encabezadoTablaCentrado),
				text.NewCol(1, "DUR.", encabezadoTablaCentrado),
				text.NewCol(1, "CURSO BÁSICO", encabezadoTablaCentrado),
				text.NewCol(3, "REQUISITOS", encabezadoTabla),
			).WithStyle(&props.Cell{BackgroundColor: gray})
		} else {
			m.AddRow(8,
				text.NewCol(3, "ASIGNATURA", encabezadoTablaCentrado),
				text.NewCol(1, "SCT", encabezadoTablaCentrado),
				text.NewCol(1, "H.P.", encabezadoTablaCentrado),
				text.NewCol(1, "H.N.P.", encabezadoTablaCentrado),
				text.NewCol(1, "TIPO", encabezadoTablaCentrado),
				text.NewCol(1, "DUR.", encabezadoTablaCentrado),
				text.NewCol(1, "CURSO BÁSICO", encabezadoTablaCentrado),
				text.NewCol(3, "REQUISITOS", encabezadoTabla),
			).WithStyle(&props.Cell{BackgroundColor: gray})
		}

		m.AddRows(text.NewRow(8, tituloGrupo, tituloTextoSemestre).WithStyle(&estiloSimbologia), row.New(5))

		for _, curso := range grupo {
			var cursobas string

			if curso.CursoBasico == 0 {
				cursobas = "No"
			} else if curso.CursoBasico == 1 {
				cursobas = "Sí"
			} else {
				cursobas = ""
			}

			requisitosTextContent := ""
			for _, requisito := range curso.MallasRequisitos {
				requisitosTextContent += " - " + requisito.MallaPionero.CursoPionero.CompetenciaAsignatura.NomCompAsig
			}

			fmt.Println("Debug requisitosTextContent:", requisitosTextContent)

			var asignaturaTextContent string
			if curso.CursoPionero.Institucional == 1 {
				asignaturaTextContent = " * " + curso.CursoPionero.CompetenciaAsignatura.NomCompAsig
			} else {
				asignaturaTextContent = curso.CursoPionero.CompetenciaAsignatura.NomCompAsig
			}

			competenciasTextContent := ""
			for _, modulo := range curso.CursoPionero.ModulosSubcompetencias {
				if len(modulo.CompetenciasRc) > 0 {
					for _, competencia := range modulo.CompetenciasRc {
						compComp := competencia.NomCompetencia
						competenciasTextContent += compComp
					}
				}
			}

			var rowHeight float64
			if idPlan == 2 {
				rowHeight = calculateRowHeight(competenciasTextContent, fontSize, 2.0, 2)
			} else {
				rowHeight = calculateRowHeight(requisitosTextContent, fontSize, 1.5, 5)
			}

			if idPlan == 2 {
				m.AddRow(rowHeight,
					col.New(0),
					text.NewCol(2, asignaturaTextContent, props.Text{Size: fontSize, Align: align.Left, Left: 1.0, Top: 1.0, Right: 1.0, VerticalPadding: 1.0}),
					text.NewCol(2, competenciasTextContent, props.Text{Size: fontSize, Align: align.Left, Left: 1.0, Top: 1.0, Right: 1.0, VerticalPadding: 1.0}),
					text.NewCol(1, strconv.FormatFloat(curso.CursoPionero.SctTotal, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, strconv.FormatFloat(curso.CursoPionero.HrsCronoPresenSemestral, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, strconv.FormatFloat(curso.CursoPionero.HrsCronoNoPresenSemestral, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, curso.TipoCursoMalla.Descripcion, props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, curso.Duracion, props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, cursobas, props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(2, requisitosTextContent, props.Text{Size: fontSize, Align: align.Left}),
				)
			} else if idPlan == 4 {
				m.AddRow(rowHeight,
					col.New(0),
					text.NewCol(3, asignaturaTextContent, props.Text{Size: fontSize, Align: align.Left, Left: 1.0, Top: 1.0, Right: 1.0, VerticalPadding: 1.0}),
					text.NewCol(1, strconv.FormatFloat(curso.CursoPionero.HrsTeoricas, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, strconv.FormatFloat(curso.CursoPionero.HrsEjercicios, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, strconv.FormatFloat(curso.CursoPionero.HrsPracticas, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, curso.TipoCursoMalla.Descripcion, props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, curso.Duracion, props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, cursobas, props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(3, requisitosTextContent, props.Text{Size: fontSize, Align: align.Left}),
				)
			} else {
				m.AddRow(rowHeight,
					col.New(0),
					text.NewCol(3, asignaturaTextContent, props.Text{Size: fontSize, Align: align.Left, Left: 1.0, Top: 1.0, Right: 1.0, VerticalPadding: 1.0}),
					text.NewCol(1, strconv.FormatFloat(curso.CursoPionero.SctTotal, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, strconv.FormatFloat(curso.CursoPionero.HrsCronoPresenSemestral, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, strconv.FormatFloat(curso.CursoPionero.HrsCronoNoPresenSemestral, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, curso.TipoCursoMalla.Descripcion, props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, curso.Duracion, props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(1, cursobas, props.Text{Size: fontSize, Align: align.Center}),
					text.NewCol(3, requisitosTextContent, props.Text{Size: fontSize, Align: align.Left}),
				)

			}

		}

	}

	//Cursos complementarios
	cursosComplementarios, err := GetMallaCursosComplementariosService(idMalla)
	if err != nil {
		return nil
	}

	//if cursosComplementarios not empty, show RAMOS COMPLEMENTARIOS
	if len(cursosComplementarios) > 0 {
		m.AddRow(8,
			text.NewCol(12, "RAMOS COMPLEMENTARIOS(CULTURALES/DEPORTIVOS ELECTIVOS E INGLÉS)", encabezadoTablaCentrado),
		).WithStyle(&props.Cell{BackgroundColor: gray})

	}

	m.AddRow(5)

	for _, cursosc := range cursosComplementarios {
		var cursobas string

		if cursosc.CursoBasico == 0 {
			cursobas = "No"
		} else if cursosc.CursoBasico == 1 {
			cursobas = "Sí"
		} else {
			cursobas = ""
		}

		requisitoscTextContent := ""
		for _, requisito := range cursosc.MallasRequisitos {
			requisitoscTextContent += requisito.MallaPionero.CursoPionero.CompetenciaAsignatura.NomCompAsig
		}

		var asignaturaTextContent string
		if cursosc.CursoPionero.Institucional == 1 {
			asignaturaTextContent = " * " + cursosc.CursoPionero.CompetenciaAsignatura.NomCompAsig
		} else {
			asignaturaTextContent = cursosc.CursoPionero.CompetenciaAsignatura.NomCompAsig
		}

		competenciascTextContent := ""
		for _, modulo := range cursosc.CursoPionero.ModulosSubcompetencias {
			if len(modulo.CompetenciasRc) > 0 {
				for _, competencia := range modulo.CompetenciasRc {
					compComp := competencia.NomCompetencia
					competenciascTextContent += compComp
				}
			}
		}

		var rowHeight float64
		if idPlan == 2 {
			rowHeight = calculateRowHeight(competenciascTextContent, fontSize, 2.0, 2)
		} else {
			rowHeight = calculateRowHeight(requisitoscTextContent, fontSize, 1.5, 5)
		}

		//Ordena alfabéticamente
		if idPlan == 2 {
			m.AddRow(rowHeight,
				col.New(0),
				text.NewCol(2, asignaturaTextContent, props.Text{Size: fontSize, Align: align.Left, Left: 1.0, Top: 1.0, Right: 1.0, VerticalPadding: 1.0}),
				text.NewCol(2, competenciascTextContent, props.Text{Size: fontSize, Align: align.Left, Left: 1.0, Top: 1.0, Right: 1.0, VerticalPadding: 1.0}),
				text.NewCol(1, strconv.FormatFloat(cursosc.CursoPionero.SctTotal, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, strconv.FormatFloat(cursosc.CursoPionero.HrsCronoPresenSemestral, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, strconv.FormatFloat(cursosc.CursoPionero.HrsCronoNoPresenSemestral, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, cursosc.TipoCursoMalla.Descripcion, props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, cursosc.Duracion, props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, cursobas, props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(2, requisitoscTextContent, props.Text{Size: fontSize, Align: align.Left}),
			)
		} else if idPlan == 4 {
			m.AddRow(rowHeight,
				col.New(0),
				text.NewCol(3, asignaturaTextContent, props.Text{Size: fontSize, Align: align.Left, Left: 1.0, Top: 1.0, Right: 1.0, VerticalPadding: 1.0}),
				text.NewCol(1, strconv.FormatFloat(cursosc.CursoPionero.HrsTeoricas, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, strconv.FormatFloat(cursosc.CursoPionero.HrsEjercicios, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, strconv.FormatFloat(cursosc.CursoPionero.HrsPracticas, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, cursosc.TipoCursoMalla.Descripcion, props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, cursosc.Duracion, props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, cursobas, props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(3, requisitoscTextContent, props.Text{Size: fontSize, Align: align.Left}),
			)
		} else {
			m.AddRow(rowHeight,
				col.New(0),
				text.NewCol(3, asignaturaTextContent, props.Text{Size: fontSize, Align: align.Left, Left: 1.0, Top: 1.0, Right: 1.0, VerticalPadding: 1.0}),
				//text.NewCol(2, competenciascTextContent, props.Text{Size: fontSize, Align: align.Left, Left: 1.0, Top: 1.0, Right: 1.0, VerticalPadding: 1.0}),
				text.NewCol(1, strconv.FormatFloat(cursosc.CursoPionero.SctTotal, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, strconv.FormatFloat(cursosc.CursoPionero.HrsCronoPresenSemestral, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, strconv.FormatFloat(cursosc.CursoPionero.HrsCronoNoPresenSemestral, 'f', -1, 64), props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, cursosc.TipoCursoMalla.Descripcion, props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, cursosc.Duracion, props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(1, cursobas, props.Text{Size: fontSize, Align: align.Center}),
				text.NewCol(3, requisitoscTextContent, props.Text{Size: fontSize, Align: align.Left}),
			)
		}
	}

	return m
}

// Establecer cuántos elementos se pueden acomodar en una línea.
func calculateRowHeight(content string, fontSize float64, lineSpacingFactor float64, elementsPerLine int) float64 {
	// Contar el número de elementos en la cadena
	numElements := len(strings.Split(content, " - "))

	// Estimar el número de líneas
	numLines := (numElements + elementsPerLine - 1) / elementsPerLine

	// Calcula la altura de la fila
	rowHeight := float64(numLines) * (fontSize * lineSpacingFactor)

	return rowHeight
}

func getPageHeader() core.Row {
	return row.New(20).Add(
		image.NewFromFileCol(3, "assets/img/Logo_UMAG_Color.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)
}

func getGrayColor() *props.Color {
	return &props.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}
