import 'package:equatable/equatable.dart';
import '../../models/student_model.dart';

abstract class StudentState extends Equatable {
  @override
  List<Object?> get props => [];
}

class StudentInitial extends StudentState {}

class StudentLoading extends StudentState {}

class StudentLoaded extends StudentState {
  final List<Student> students;

  StudentLoaded(this.students);

  @override
  List<Object?> get props => [students];
}

class StudentError extends StudentState {
  final String message;

  StudentError(this.message);

  @override
  List<Object?> get props => [message];
}
