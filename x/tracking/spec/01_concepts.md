<!--
order: 1
-->

# Concepts

## The Tracking module

The Tracking module is primarily created to track and process events emitted by loan contract.


## Module requirements:
* The module implements tracking of loan contract events using EVM hooks;
* Module proceed next event types: borrowed and liquidated.
* Module stores all borrowing positions and check the ability to liquidate them in `EndBlocker`.
* After liquidation an amount of deposit in liquidated position is distributed to validators as a reward.
* If user position can be liquidated it has to be immediately liquidated by liquidator in the `EndBlocker`.
