import requests
import json, os, time

EASY = "https://opentdb.com/api.php?amount=10&category=31&difficulty=easy&type=multiple"
MEDIUM = "https://opentdb.com/api.php?amount=10&category=31&difficulty=medium&type=multiple"
HARD = "https://opentdb.com/api.php?amount=10&category=31&difficulty=hard&type=multiple"

def set_questions():
    difficulty = get_difficulty()
    if difficulty == "1" or difficulty.lower() == "easy": response = requests.get(EASY)
    elif difficulty == "2" or difficulty.lower() == "medium": response = requests.get(MEDIUM)
    elif difficulty == "3" or difficulty.lower() == "hard": response = requests.get(HARD)
    else: print("Invalid Difficulty"); exit(0)
    Questions_Dict = json.loads(response.text)

    questions = []
    correct_answer = []
    incorrect_answers = []
    for data in Questions_Dict["results"]:
        questions.append(data["question"])
        correct_answer.append(data["correct_answer"])
        incorrect_answers.append(data["incorrect_answers"])

    return list(zip(questions, correct_answer, incorrect_answers))


def get_difficulty():
    print("1 - Easy")
    print("2 - Medium")
    print("3 - Hard")
    difficulty = input("Choose Difficulty: ")
    return difficulty

def game():
    QUESTIONS = set_questions()
    SLEEP = 1.5
    points = 0

    question_count = 0
    while len(QUESTIONS) > question_count:
        os.system("cls")
        correct_answer = []
        correct_answer.append(QUESTIONS[question_count][1])

        print(QUESTIONS[question_count][0] + "\n")
        print(" ".join(n + " - " for n in correct_answer + QUESTIONS[question_count][2])[0:-2])
        your_answer = input("Enter Your Answer: ")

        if your_answer.title() == QUESTIONS[question_count][1]:
            print("Good Job! You Earn 10 Points")
            points += 10
            print("Total Points: {}".format(points))
            question_count += 1
            time.sleep(SLEEP)
        elif your_answer.title() in QUESTIONS[question_count][2]:
            print("Sorry, Wrong Answer")
            question_count += 1
            time.sleep(SLEEP)
        else:
            print("There is No Such Choice")
            time.sleep(SLEEP)

def main():
    game()
    
    

main()
