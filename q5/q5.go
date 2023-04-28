package q5
import "strings"

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	func ProcessString(s string) []string {
	var letras []string

	A := "A"
	E := "E"
	I := "I"
	O := "O"
	U := "U"

	for i := 0; i < len(letras); i++ {
		if letras[i] == A {
			strings.Split(letras[i], A)

		} else if letras[i] == E {
			strings.Split(letras[i], E)

		} else if letras[i] == I {
			strings.Split(letras[i], I)

		} else if letras[i] == O {
			strings.Split(letras[i], O)
		} else if letras[i] == U {
			strings.Split(letras[i], U)
		}
		var novaString []string
		novaString = append(novaString, letras...)
		return novaString

	}

}
	

