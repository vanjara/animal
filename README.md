# animal

# What is the game
# User selects an animal which the game has to guess
# Game asks user yes/no questions to guess the answer
# Game should be able to remember information of prior runs
# Ex:
# User decides zebra
# Game can ask questions like:
# Game: Is it an elephant (yes/no)?
# User: no
# Game: Is it a horse (yes/no)?
# User: no
# Game: Does the animal have stripes (yes/no)?
# User: yes
# Game: It is a giraffe (yes/no)?
# User: no
# Game: Is it a zebra (yes/no)?
# User: yes
# Game: Thank you for playing the game! Would you like to play again (yes/no)?
# If yes, repeat above
# If not, exit
# Other notes:
# How does game know about which animals to guess
# What would be a question bank for the game
# How will the game remember answers from previous attempts/game runs

# Adding a new animal
# Is it a Zebra
# ans is No (You stumped me)
# Please tell me the animal you were thinking about?
# Suppose Ans is Tiger
# What would distinguish zebra from tiger
# need a yes/no question
# Is it a predator?
# We should ask this question to the user for a tiger?
# Does it have 4 legs?
# Does it have stripes?
# Is it a predator?
# Ans is Tiger

# Add the new question to the map
# Add the new animal to the map
# connect the yes/no answer to the new animal
# other answer will be AnswerLose
# Replace the AnswerLose with the question that got us here