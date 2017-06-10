package main

import (
	"fmt"

	. "bayesian"
)

//https://gowalker.org/github.com/jbrukh/bayesian
func main() {

	classifier()

}

func classifier() {

	// I - Definição das classes
	const (
		Positive Class = "Positive" /* 0 */
		Negative Class = "Negative" /* 1 */
	)
	//classifier := NewClassifierTfIdf(Positive, Negative)
	classifier := NewClassifier(Positive, Negative)

	// II - Treinamento & III - Aprendizado
	classifier.Learn([]string{"Dog"}, Positive)
	classifier.Learn([]string{"Love"}, Positive)
	classifier.Learn([]string{"Love"}, Positive)
	classifier.Learn([]string{"Dog"}, Positive)
	classifier.Learn([]string{"House"}, Positive)
	classifier.Learn([]string{"Cat"}, Positive)
	classifier.Learn([]string{"House"}, Positive)
	classifier.Learn([]string{"Love"}, Positive)
	// classifier.Learn([]string{"Mike", "have", "a", "Dog"}, Positive)
	// classifier.Learn([]string{"We", "Love", "golang"}, Positive)
	// classifier.Learn([]string{"I", "Love", "locaweb"}, Positive)
	// classifier.Learn([]string{"I", "have", "a", "Dog"}, Positive)
	// classifier.Learn([]string{"It's", "my", "House", "come", "on"}, Positive)
	// classifier.Learn([]string{"I", "have", "a", "Cat"}, Positive)
	// classifier.Learn([]string{"I", "have", "a", "big", "House"}, Positive)
	// classifier.Learn([]string{"I", "Love", "this", "song"}, Positive)

	classifier.Learn([]string{"Love"}, Negative)
	classifier.Learn([]string{"Bad"}, Negative)
	classifier.Learn([]string{"Cat"}, Negative)
	classifier.Learn([]string{"Cat"}, Negative)
	classifier.Learn([]string{"Love"}, Negative)
	// classifier.Learn([]string{"I", "don't", "Love", "you"}, Negative)
	// classifier.Learn([]string{"I'm", "feeling", "so", "Bad"}, Negative)
	// classifier.Learn([]string{"I", "hate", "my", "Cat"}, Negative)
	// classifier.Learn([]string{"This", "Cat", "hurt", "me"}, Negative)
	// classifier.Learn([]string{"I", "don't", "Love", "this", "song"}, Negative)
	//classifier.ConvertTermsFreqToTfIdf()

	// IV - Coleta de Dados & V - Split dos atributos a serem classificados
	atribute := []string{"Dog"}
	// VI - Classificação
	scores, likely, _ := classifier.ProbScores(atribute)
	fmt.Printf("%s - %7.5f %v\n", atribute, scores[likely], classifier.Classes[likely])

	atribute = []string{"Love"}
	scores, likely, _ = classifier.ProbScores(atribute)
	fmt.Printf("%s - %7.5f %v\n", atribute, scores[likely], classifier.Classes[likely])

	atribute = []string{"Bad"}
	scores, likely, _ = classifier.ProbScores(atribute)
	fmt.Printf("%s - %7.5f %v\n", atribute, scores[likely], classifier.Classes[likely])

	atribute = []string{"House"}
	scores, likely, _ = classifier.ProbScores(atribute)
	fmt.Printf("%s - %7.5f %v\n", atribute, scores[likely], classifier.Classes[likely])

	atribute = []string{"Cat"}
	scores, likely, _ = classifier.ProbScores(atribute)
	fmt.Printf("%s - %7.5f %v\n", atribute, scores[likely], classifier.Classes[likely])

}
