import 'package:equatable/equatable.dart';
import '../../models/stand_model.dart';

abstract class StandState extends Equatable {
  @override
  List<Object> get props => [];
}

class StandInitial extends StandState {}

class StandLoading extends StandState {}

class StandLoaded extends StandState {
  final List<Stand> stands;

  StandLoaded(this.stands);

  @override
  List<Object> get props => [stands];
}

class StandError extends StandState {
  final String message;

  StandError(this.message);

  @override
  List<Object> get props => [message];
}
