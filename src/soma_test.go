package main

func TestSoma(t *testing.T){
  resultado := Soma(5,5)
  if resultado != 10{
    t.Errorf("Soma esperada: %d, obtida: %d", 10, resultado)
  }
}