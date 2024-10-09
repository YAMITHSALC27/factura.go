package main

import (
	"fmt"
	"strconv"
	"time"
	"github.com/jung-kurt/gofpdf"
)

type Producto struct {
	nombre     string
	cantidad   int
	precioUnit float64
	total      float64
}

func main() {
	//solicitar informacion del cliente
	var nombreCliente, fechaFactura string
	fmt.Println("porfavor ingrese el nombre del cliente: ")
	fmt.Scanln(&nombreCliente)

	var fechaValida bool
	for !fechaValida {
		fmt.Println("porfavor ingrese la fecha de la factura (aaaa/mm/dd): ")
		fmt.Scanln(&fechaFactura)

		//convertir_fecha_ingresada_a_tipo_time
		_, err := time.Parse("2006/01/02", fechaFactura)
		if err != nil {
			fmt.Println("fecha no valida, por favor ingrese un formato valido (ejemplo: 2024/01/01): ")
		} else {
			fechaValida = true
		}
	}

	//solicitar un producto
	var (
		productos      []Producto
		nombreProducto string
		cantidad       int
		precioUnit     float64
	)

	for {
		fmt.Println("ingrese el nombre del producto (o 'fin' para terminar): ")
		fmt.Scanln(&nombreProducto)
		if nombreProducto == "fin" {
			break
		}

		fmt.Println("ingrese la cantidad del producto: ")
		fmt.Scanln(&cantidad)

		fmt.Println("ingrese el precio unitario: ")
		fmt.Scanln(&precioUnit)

		totalProducto := float64(cantidad) * precioUnit
		producto := Producto{
			nombre:     nombreProducto,
			cantidad:   cantidad,
			precioUnit: precioUnit,
			total:      totalProducto,
		}

		productos = append(productos, producto)
	}

	//calcular el total del cliente
	var totalCliente float64
	for _, producto := range productos {
		totalCliente += producto.total
	}

	//mostrar resumen de factura
	fmt.Println("\n------Resumen de la factura------")
	fmt.Printf("cliente: %s\n", nombreCliente)
	fmt.Printf("fecha: %s\n", fechaFactura)
	fmt.Println("productos: ")
	for _, producto := range productos {
		fmt.Printf("- %s: %d x %.2f = %.2f\n", producto.nombre, producto.cantidad, producto.precioUnit, producto.total)
	}
	fmt.Printf("Total General: %.2f\n", totalCliente)
	fmt.Println("----------------------------------")

	var generarPDF string
	fmt.Println("Desea generar un PDF de la factura? (s/n): ")
	fmt.Scanln(&generarPDF)

	if generarPDF == "s" || generarPDF == "S" {
		//crear PDF con gofpdf
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

		//titulo de la factura
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(190, 10, "FACTURA DE VENTAS")
		pdf.Ln(12)

		//agregar nombre cliente y fecha
		pdf.SetFont("Arial", "", 12)
		pdf.Cell(40, 10, fmt.Sprintf("CLIENTE: %s", nombreCliente))
		pdf.Ln(10)
		pdf.Cell(40, 10, fmt.Sprintf("FECHA: %s", fechaFactura))
		pdf.Ln(10)

		pdf.SetFillColor(200, 200, 200)
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(60, 10, "PRODUCTO", "1", 0, "C", true, 0, "")
		pdf.CellFormat(30, 10, "CANTIDAD", "1", 0, "C", true, 0, "")
		pdf.CellFormat(40, 10, "PRECIO UNITARIO", "1", 0, "C", true, 0, "")
		pdf.CellFormat(40, 10, "TOTAL", "1", 1, "C", true, 0, "")

		//agregar productos
		pdf.SetFont("Arial", "", 12)
		for _, producto := range productos {
			pdf.CellFormat(60, 10, producto.nombre, "1", 0, "L", false, 0, "")
			pdf.CellFormat(30, 10, strconv.Itoa(producto.cantidad), "1", 0, "C", false, 0, "")
			pdf.CellFormat(40, 10, fmt.Sprintf("%.2f", producto.precioUnit), "1", 0, "R", false, 0, "")
			pdf.CellFormat(40, 10, fmt.Sprintf("%.2f", producto.total), "1", 1, "R", false, 0, "")
		}

		//agregar total general
		pdf.Ln(10)
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(90, 9, "")
		pdf.CellFormat(40, 10, "TOTAL GENERAL", "1", 0, "R", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("%.2f", totalCliente), "1", 1, "R", false, 0, "")

		//guardar y mostrar PDF en el navegador
		err := pdf.OutputFileAndClose("factura.pdf")
		if err != nil {
			fmt.Printf("Error al generar el PDF: %s", err.Error())
		}

		fmt.Println("PDF generado correctamente en la ruta: 'factura.pdf'")

	} else {
		fmt.Println("No se generó el PDF.")
	}

}
