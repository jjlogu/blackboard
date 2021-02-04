import random

#TODO
#create a cell class
#handle exceptions
# what if player doesn't have money?

class empty:
    def action(self, player):
        pass

class jail:
    def action(self, player):
        if player.money >= 150:
            player.money -= 150

class treasure:
    def action(self, player):
        if player.money >= 200:
            player.money += 200

class hotel:
    def __init__(self, worth=200):
        self.owner = None
        self.worth = worth

    def action(self, player):
        if self.owner:
            self.pay_rent(player)
        else:
            if player.money >=200:
                player.money -= 200
                self.owner = player
    def pay_rent(self, player):
        player.money -= 50
        self.owner.money += 50


class dice:
    def __init__(self):
        pass
    def roll(self):
        # return a random number between 2-12
        return random.randrange(2,12)

class board:
    def __init__(self, cell_posi_type):
        self.cell_posi_type = cell_posi_type
        self.max_position = len(cell_posi_type)
    def visit_cell(self, position, player):
        self.cell_posi_type[position].action(player)

class player:
    def __init__(self, name, money):
        self.name = name
        self.money = money

class game:
    def __init__(self, num_players, initial_money, board, dice, max_chance):
        self.players = []
        for i in range(num_players):
            self.players.append(player("Player{}".format(i), initial_money))
        self.board = board
        self.dice = dice
        self.max_chance = max_chance

    def play(self):
        # TODO: break down further
        for _ in range(self.max_chance):
            for player in self.players:
                position = self.dice.roll()
                self.board.visit_cell(position, player)

class resultDisplayer:
    def __init__(self, gam):
        self.game = gam

    def display(self):
        self.game.players.sort(key=lambda player:player.money, reverse=True)
        for player in self.game.players:
            print(player.name, ":", player.money)

# Get the input

if __name__ == "__main__":
    num_players = int(input())
    cell_posi_type = []
    for cell in ['E', 'E', 'J', 'H', 'E', 'T', 'J', 'T', 'E', 'E', 'H', 'J', 'T', 'H', 'E', 'E', 'J', 'H', 'E', 'T', 'J', 'T', 'E', 'E', 'H', 'J', 'T', 'H', 'J', 'E', 'E', 'J', 'H', 'E', 'T', 'J', 'T', 'E', 'E', 'H', 'J', 'T', 'E', 'H', 'E']:
        if cell == 'E':
            cell_posi_type.append(empty())
        elif cell == 'J':
            cell_posi_type.append(jail())
        elif cell == 'T':
            cell_posi_type.append(treasure())
        elif cell == 'H':
            cell_posi_type.append(hotel())
        else:
           NotImplemented("Cell type")

    bord = board(cell_posi_type)
    dic = dice()


    gam = game(num_players, 1000, bord, dic, 10)

    gam.play()


    result = resultDisplayer(gam)
    result.display()


