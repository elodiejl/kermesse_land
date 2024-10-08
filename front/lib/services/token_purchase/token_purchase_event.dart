import 'package:equatable/equatable.dart';

abstract class TokenPurchaseEvent extends Equatable {
  @override
  List<Object> get props => [];
}

class PurchaseTokens extends TokenPurchaseEvent {
  final String token;
  final String parentId;
  final int tokenAmount;

  PurchaseTokens(this.token, this.parentId, this.tokenAmount);

  @override
  List<Object> get props => [parentId, tokenAmount];
}