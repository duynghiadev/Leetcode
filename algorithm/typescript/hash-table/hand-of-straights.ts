// Source: https://leetcode.com/problems/hand-of-straights/description/

function isNStraightHand(hand: number[], groupSize: number): boolean {
  if (hand.length % groupSize !== 0) {
    return false; // If the total cards can't be evenly divided into groups, return false
  }

  // Count the frequency of each card
  const cardCount = new Map<number, number>();
  for (const card of hand) {
    cardCount.set(card, (cardCount.get(card) || 0) + 1);
  }

  // Sort the cards to process the smallest card first
  const sortedCards = Array.from(cardCount.keys()).sort((a, b) => a - b);

  for (const card of sortedCards) {
    const count = cardCount.get(card) || 0;

    // If the card count is greater than 0, form a group starting with this card
    if (count > 0) {
      for (let i = 0; i < groupSize; i++) {
        const currentCardCount = cardCount.get(card + i) || 0;

        if (currentCardCount < count) {
          return false; // Not enough cards to form a valid group
        }

        cardCount.set(card + i, currentCardCount - count); // Reduce the count for the current card
      }
    }
  }

  return true;
}

console.log(isNStraightHand([1, 2, 3, 6, 2, 3, 4, 7, 8], 3));
console.log(isNStraightHand([1, 2, 3, 4, 5], 4));
console.log(isNStraightHand([8, 10, 12], 3));
console.log(isNStraightHand([1, 2, 3, 4, 5, 6], 2));
