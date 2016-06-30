package main

import "testing"

func TestFuncionSuma(t *testing.T){

   resultado := suma(2,3)
   if resultado != 5 {
     t.Errorf("La suma no corresponde a lo esperado obtenido: %d esperado: %d",resultado, 5)
   }
}
