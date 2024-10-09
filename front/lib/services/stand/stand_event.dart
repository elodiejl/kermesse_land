import 'package:equatable/equatable.dart';

abstract class StandEvent extends Equatable {
  @override
  List<Object> get props => [];
}

class FetchStands extends StandEvent {
  final String token;

  FetchStands(this.token);

  @override
  List<Object> get props => [token];
}

class FetchStandsByKermesse extends StandEvent {
  final String token;
  final int kermesseId;

  FetchStandsByKermesse(this.token, this.kermesseId);

  @override
  List<Object> get props => [token, kermesseId];
}

