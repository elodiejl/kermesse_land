import 'package:equatable/equatable.dart';

abstract class StudentEvent extends Equatable {
  @override
  List<Object?> get props => [];
}

class FetchStudents extends StudentEvent {
  final int parentId;
  final String token;

  FetchStudents(this.parentId, this.token);

  @override
  List<Object?> get props => [parentId, token];
}
