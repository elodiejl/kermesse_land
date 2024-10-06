
import 'package:equatable/equatable.dart';

abstract class TokenPurchaseState extends Equatable {
  @override
  List<Object> get props => [];
}

class TokenPurchaseInitial extends TokenPurchaseState {}

class TokenPurchaseLoading extends TokenPurchaseState {}

class TokenPurchaseSuccess extends TokenPurchaseState {
  final String transactionId;

  TokenPurchaseSuccess(this.transactionId);

  @override
  List<Object> get props => [transactionId];
}

class TokenPurchaseError extends TokenPurchaseState {
  final String error;

  TokenPurchaseError(this.error);

  @override
  List<Object> get props => [error];
}
