"""
Internal Representation of Poker Hands
Note that our wrapper will convert strings such as "4h" to Card(Rank.four, Suit.heart)
"""

from enum import Enum

# Suit
# Enumeration of card suits
class Suit (Enum):
  heart   = 1
  club    = 2
  diamond = 3
  spade   = 4

# Rank
# Enumeration of card ranks
class Rank (Enum):
  two   = 2
  three = 3
  four  = 4
  five  = 5
  six   = 6
  seven = 7
  eight = 8
  nine  = 9
  ten   = 10
  jack  = 11
  queen = 12
  king  = 13
  ace   = 14

# Hand
# Enumeration of poker hands
class Hand (Enum):
  junk      = 1
  pair      = 2
  twoPair   = 3
  triple    = 4
  straight  = 5
  flush     = 6
  fullHouse = 7


# Card
# Object representing a single card
class Card:
  def __init__ (self, rank, suit):
    self.rank = rank
    self.suit = suit

  def __lt__ (self, rhs):
    return self.rank < rhs.rank

# Hand
# Object representing a collection of 5 cards
class Hand:
  def __init (self, *cards):
    if cards.length != 5: raise "wtf"
    self.cards = Sorted(cards)

  def high_card (self):
    return self.cards[-1]