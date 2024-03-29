# Guess the animal game

## What is the game?
User selects an animal which the game has to guess
Game asks user yes/no questions to guess the answer
Game should be able to remember information of prior runs
For Ex:
User decides zebra
Game can ask questions like:
```
Game: Is it an elephant (yes/no)?
User: no
Game: Is it a horse (yes/no)?
User: no
Game: Does the animal have stripes (yes/no)?
User: yes
Game: It is a giraffe (yes/no)?
User: no
Game: Is it a zebra (yes/no)?
User: yes
Game: Thank you for playing the game! Would you like to play again (yes/no)?*
```
If yes, repeat above
If not, exit
## Other notes:
- How does game know about which animals to guess
- What would be a question bank for the game
- How will the game remember answers from previous attempts/game runs

## Adding a new animal
```
Is it a Zebra
ans is No (You stumped me)
Please tell me the animal you were thinking about?
Suppose Ans is Tiger
What would distinguish zebra from tiger
need a yes/no question
Is it a predator?
We should ask this question to the user for a tiger?
Does it have 4 legs?
Does it have stripes?
Is it a predator?
Ans is Tiger
```
- Add the new question to the map - Done 
- Add the new animal to the map - Done
- connect the yes/no answer to the new animal - Done
- Other answer will be AnswerLose
- Replace the AnswerLose with the question that got us here - Done
- Add a Play Again Feature - Done
- Add the new question to the PlayAgainFeature - Done

### StartingData example
```
var StartingData = map[string]Question{
	"Does it have 4 legs?": Question{
		Yes: "Does it have stripes?",
		No:  "Is it carnivorous?",
	},
	"Does it have stripes?": Question{
		Yes: "Is it a zebra?",
		No:  "Is it a lion?",
	},
    "Does it have stripes?": Question{
        Yes: "Is it a predator?",
        No: "Is it a lion?"
    }
    "Is it a predator?": Question{
        Yes: "Is it a tiger?",
        No: "Is it a zebra?"
    }
	"Is it carnivorous?": Question{
	Yes: "Is it a snake?",
	No:  "Is it a worm?",
	},
	"Is it a zebra?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},	
    "Is it a tiger?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},
	"Is it a giraffe?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},
	"Is it a lion?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},
	"Is it a snake?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},
	"Is it a worm?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},
```
- Read Starting Data from a file (any format) - Done
- Then write new data to the file for future runs
- Move replay logic to the game (not in Main) - Done
- GetUserYesorNo
- Loop until valid answers (no need to return error) - Done
- Game should use a pointer (we are modifying the game and data on learning) - Done

### Future versions 
- Play it in a browser?
- Auto deployment somewhere?
- Multi player support?
- How to write back/learn new animal? 
- Write back once a game finishes (multiple others could be running)
- Database backend?
- How much can I reuse if it was a general guessing game
- General library for any guessing game
- Could there be other ways of teaching the game (wikipedia or something else or some format database)
- What if we had a csv/data file and learn from it