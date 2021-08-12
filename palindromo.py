def is_palindromo(text):
    text = text.casefold().replace(" ","")
    text_reverse = text[::-1]
    if text == text_reverse:
        print('Es palindromo')
    else:
        print('No es palindromo')
        
is_palindromo(" Anita  lava la tina ")
is_palindromo(" Ansdfs ")
is_palindromo("amA ")
is_palindromo("hola")
