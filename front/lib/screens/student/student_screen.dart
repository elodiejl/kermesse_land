import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../services/student/student_bloc.dart';
import '../../services/student/student_event.dart';
import '../../services/student/student_state.dart';

class StudentsScreen extends StatelessWidget {
  final int parentId;
  final String token;

  const StudentsScreen({super.key, required this.parentId, required this.token});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Enfants')),
      body: BlocProvider(
        create: (context) => StudentBloc()..add(FetchStudents(parentId, token)),
        child: const StudentList(),
      ),
    );
  }
}

class StudentList extends StatelessWidget {
  const StudentList({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<StudentBloc, StudentState>(
      builder: (context, state) {
        if (state is StudentLoading) {
          return const Center(child: CircularProgressIndicator());
        } else if (state is StudentLoaded) {
          return ListView.builder(
            itemCount: state.students.length,
            itemBuilder: (context, index) {
              final student = state.students[index];
              return ListTile(
                title: Text(student.name),
                subtitle: Text('ID: ${student.id}'),
              );
            },
          );
        } else if (state is StudentError) {
          return Center(child: Text(state.message));
        }
        return const Center(child: Text('Aucun enfant trouvé.'));
      },
    );
  }
}
