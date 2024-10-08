import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../models/student_model.dart';
import '../../utils/config_io.dart';
import 'student_event.dart';
import 'student_state.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class StudentBloc extends Bloc<StudentEvent, StudentState> {
  StudentBloc() : super(StudentInitial()) {
    on<FetchStudents>(_onFetchStudents);
  }

  Future<void> _onFetchStudents(FetchStudents event, Emitter<StudentState> emit) async {
    emit(StudentLoading());
    try {
      final response = await http.get(
        Uri.parse('${Config.baseUrl}/students/parent/${event.parentId}'),
        headers: {'Authorization': 'Bearer ${event.token}'},
      );

      if (response.statusCode == 200) {
        if (kDebugMode) {
          print(response.body);
        }
        List<Student> students = (json.decode(response.body) as List)
            .map((data) => Student.fromJson(data))
            .toList();

        emit(StudentLoaded(students));
      } else {
        emit(StudentError('Failed to load students'));
      }
    } catch (e) {
      emit(StudentError('An error occurred: $e'));
    }
  }
}
