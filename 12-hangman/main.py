import random
import os

with open("words.txt", "r") as f:
    words = f.read().lower().split("\n")


# Print The Board
def print_board(the_man, found):
    os.system("cls")
    print(" |-------")
    print(" |      |")
    print(" |      {}".format(the_man[0]))
    print(" |     {}{}{}".format(the_man[1], the_man[2], the_man[3]))
    print(" |     {} {}".format(the_man[4], the_man[5]))
    print("---")

    output = ""
    for letter in found:
        output += letter

    print(output)
    return output

def main():
    the_man = ["o", "-", "|", "-", "/", "\\"]
    the_mans_place = ["", "", "", "", "", ""]

    the_word_str = random.choices(words)[0]
    the_word = list(the_word_str)
    wrong_guess_count = 0

    found = ["-" for _ in range(len(the_word))]
    visited = set()
    game_over = False
    while not game_over:
        out = print_board(the_mans_place, found)

        # WIN CHECK
        if the_word_str == out:
            print("You Win!")
            game_over = True
            break

        # CHECH THE GUESS
        guess = input("Enter Your Guess: ")
        
        if guess in the_word:
            if the_word.count(guess) == 1:
                found[the_word.index(guess)] = guess
            else:
                for _ in range(the_word.count(guess)):
                    found[the_word.index(guess)] = guess
                    the_word[the_word.index(guess)] = None
        else:
            if guess not in found and guess not in visited:
                wrong_guess_count += 1
                the_mans_place[wrong_guess_count - 1] = the_man[wrong_guess_count - 1]
            else:
                print("You already type that.")

        # GAME OVER CHECK
        if wrong_guess_count == len(the_mans_place):
            print_board(the_mans_place, found)
            print("You Lost")
            game_over = True
    
        visited.add(guess)

main()
