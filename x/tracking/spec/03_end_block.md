# EndBlock: Position Liquidation

## Overview

The `EndBlocker` function is responsible to checking the ability to liquidate position, liquidation and passing liquidation to distribution.

Position Liquidation Information:
- **From** address of liquidator which emmit a  `contactCallLiquidated`
- **Contract** address of loan contract.
- **PositionAddress** unique position address.
- **Recipient** liquidator address who is the recipient of deposit in liquidated position.

## Purpose of EndBlocker

The `EndBlocker` iterates through positions checking whether can they be liquidated or not and if yes liquidates them.

### Key Steps

1. **Retrieve positions**: The function retrieves all positions.
2. **Update Oracle price**: Update oracle price by calling the oracle contract.
3. **Check possibility of position liquidation**: It iterates through positions calling `contactCallIsUserCanBeLiquidatedMethod` contract method for each, checking if the position can be liquidated.
4. **Liquidate position**: If position can be liquidated it calls a `contactCallLiquidateMethod` passing an address of position.
5. **Calculating total amount of liquidations**: After calling `contactCallLiquidateMethod` it parses `Liquidated` events from receipt to obtain an amount of liquidation.
6. **Withdrawing liquidations as a reward** It sends an amount of liquidated tokens from liquidator account to `fee_collector` module account.