import '../../models/parent_model.dart';

abstract class ParentState {}

class ParentInitial extends ParentState {}

class ParentLoading extends ParentState {}

class ParentSuccess extends ParentState {
  final Parent parent;

  ParentSuccess(this.parent);
}

class ParentError extends ParentState {
  final String message;

  ParentError(this.message);
}
