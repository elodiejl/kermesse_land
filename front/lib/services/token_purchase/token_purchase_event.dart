import 'package:equatable/equatable.dart';

abstract class TokenPurchaseEvent extends Equatable {
  @override
  List<Object> get props => [];
}

class PurchaseTokens extends TokenPurchaseEvent {
  final String parentId;
  final int tokenAmount;

  PurchaseTokens(this.parentId, this.tokenAmount);

  @override
  List<Object> get props => [parentId, tokenAmount];
}