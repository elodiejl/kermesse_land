abstract class TransactionEvent {}

class CreateTransactionEvent extends TransactionEvent {
  final int parentId;
  final int price;
  final int tokensAmount;

  CreateTransactionEvent({
    required this.parentId,
    required this.price,
    required this.tokensAmount,
  });
}

class FetchTransactionEvent extends TransactionEvent {
  final int transactionId;

  FetchTransactionEvent(this.transactionId);
}

class FetchTransactionsByParentIdEvent extends TransactionEvent {
  final String parentId;

  FetchTransactionsByParentIdEvent(this.parentId);
}
