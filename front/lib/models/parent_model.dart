class Parent {
  final int userId;
  final int tokensAmount;

  Parent({
    required this.userId,
    required this.tokensAmount,
  });

  factory Parent.fromJson(Map<String, dynamic> json) {
    return Parent(
      userId: json['user_id'],
      tokensAmount: json['tokens_amount_available'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'user_id': userId,
      'tokens_amount_available': tokensAmount,
    };
  }
}
