

#[echo "What's yout name?" 
var name: string = readline(stdin)
echo "Hi, ", name, "!"]#

type Answer = enum aYes, aNo

proc ask(question: string): Answer =
  echo(question, " (y/n)")
  while true:
    case readLine(stdin)
    of "y", "Y", "yes", "Yes":
      return Answer.aYes  # Enums can be qualified
    of "n", "N", "no", "No":
      return Answer.aNo
    else: echo("Please be clear: yes or no")

proc addSugar(amount: int = 2) = # Default amount is 2, returns nothing
  assert(amount > 0 and amount < 9000, "Crazy Sugar")
  for a in 1..amount:
    echo(a, " sugar...")

case ask("Would you like sugar in your tea?")
of aYes:
  addSugar(3)
of aNo:
  echo "Oh do take a little!"
  addSugar()