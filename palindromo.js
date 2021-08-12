const isPalindromo = (text) => {
  text = text.toLowerCase().replace(/ /g, "")
  let textReverse = text.split("").reverse().join("")
  text === textReverse
    ? console.log("Es Palindromo")
    : console.log("No es Palindromo")
}

isPalindromo(" Anita  lava la tina ")
isPalindromo(" Ansdfs ")
isPalindromo("amA ")
isPalindromo("hola")
