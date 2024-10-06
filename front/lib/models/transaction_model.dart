class Transaction {
  final int id;
  final int parentId;
  final int price;
  final int tokensAmount;
  final String transactionDate;

  Transaction({
    required this.id,
    required this.parentId,
    required this.price,
    required this.tokensAmount,
    required this.transactionDate,
  });

  factory Transaction.fromJson(Map<String, dynamic> json) {
    return Transaction(
      id: json['id'],
      parentId: json['parent_id'],
      price: json['price'],
      tokensAmount: json['tokens_amount'],
      transactionDate: json['transaction_date'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'parent_id': parentId,
      'price': price,
      'tokens_amount': tokensAmount,
      'transaction_date': transactionDate,
    };
  }
}
