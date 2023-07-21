package keeper_test

// func TestCalculateRewards(t *testing.T) {
// 	keeper, _ := keepertest.PoaKeeper(t)

// 	// Set up test data
// 	totalAmount := sdk.NewDec(1000)
// 	scores := []sdk.Dec{
// 		sdk.NewDec(50),
// 		sdk.NewDec(30),
// 		sdk.NewDec(20),
// 	}

// 	// Execute the function to be tested
// 	rewards := keeper.CalculateRewards(totalAmount, scores)

// 	// Validate the results

// 	// Check if the number of rewards matches the number of scores
// 	require.Len(t, rewards, len(scores), "Incorrect number of rewards")

// 	// Check if the sum of rewards equals the total amount
// 	sumRewards := sdk.ZeroDec()
// 	for _, reward := range rewards {
// 		sumRewards = sumRewards.Add(reward)
// 	}
// 	require.True(t, totalAmount.Equal(sumRewards), "Sum of rewards does not match totalAmount")

// 	// Check if each reward is correctly calculated
// 	for i, score := range scores {
// 		expectedReward := totalAmount.Mul(score.Quo(calculateTotalScore(scores)))
// 		require.True(t, expectedReward.Equal(rewards[i]), "Incorrect reward calculation for score %s", score.String())
// 		t.Logf("Score: %s, Expected Reward: %s, Actual Reward: %s", score.String(), expectedReward.String(), rewards[i].String())
// 	}
// }

// // Helper function to calculate the total score for test validation
// func calculateTotalScore(scores []sdk.Dec) sdk.Dec {
// 	totalScore := sdk.ZeroDec()
// 	for _, score := range scores {
// 		totalScore = totalScore.Add(score)
// 	}
// 	return totalScore
// }
