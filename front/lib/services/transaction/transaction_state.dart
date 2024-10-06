import '../../models/transaction_model.dart';

abstract class TransactionState {}

class TransactionInitial extends TransactionState {}

class TransactionLoading extends TransactionState {}

class TransactionSuccess extends TransactionState {
  final Transaction transaction;

  TransactionSuccess(this.transaction);
}

class TransactionsSuccess extends TransactionState {
  final List<Transaction> transactions;

  TransactionsSuccess(this.transactions);
}

class TransactionError extends TransactionState {
  final String message;

  TransactionError(this.message);
}
