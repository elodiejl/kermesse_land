import 'package:bloc/bloc.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/models/parent_model.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'parent_event.dart';
import 'parent_state.dart';
import '../../utils/config_io.dart';

class ParentBloc extends Bloc<ParentEvent, ParentState> {
  ParentBloc() : super(ParentInitial()) {
    on<CreateParent>(_onCreateParent);
    on<GetParentById>(_onGetParentById);
    on<UpdateParent>(_onUpdateParent);
    on<DeleteParent>(_onDeleteParent);
  }

  // Gestionnaire pour l'événement CreateParent
  void _onCreateParent(CreateParent event, Emitter<ParentState> emit) async {
    emit(ParentLoading());
    try {
      final response = await http.post(
        Uri.parse('${Config.baseUrl}/parents'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'user_id': event.userId,
          'tokens_amount_available': event.tokensAmount,
        }),
      );

      if (response.statusCode == 201) {
        final parentData = jsonDecode(response.body);
        emit(ParentSuccess(parentData));
      } else {
        emit(ParentError('Failed to create parent'));
      }
    } catch (e) {
      emit(ParentError(e.toString()));
    }
  }

  void _onGetParentById(GetParentById event, Emitter<ParentState> emit) async {
    emit(ParentLoading());
    try {
      final response = await http.get(
        Uri.parse('${Config.baseUrl}/parents/${event.id}'),
      );

      if (response.statusCode == 200) {
        final parentData = Parent.fromJson(jsonDecode(response.body));
        emit(ParentSuccess(parentData));
      } else {
        emit(ParentError('Parent not found'));
      }
    } catch (e) {
      emit(ParentError(e.toString()));
    }
  }


  // Gestionnaire pour l'événement UpdateParent
  void _onUpdateParent(UpdateParent event, Emitter<ParentState> emit) async {
    emit(ParentLoading());
    try {
      final response = await http.put(
        Uri.parse('${Config.baseUrl}/parents/${event.id}'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'tokens_amount_available': event.tokensAmount,
        }),
      );

      if (response.statusCode == 200) {
        final parentData = jsonDecode(response.body);
        emit(ParentSuccess(parentData));
      } else {
        emit(ParentError('Failed to update parent'));
      }
    } catch (e) {
      emit(ParentError(e.toString()));
    }
  }

  // Gestionnaire pour l'événement DeleteParent
  void _onDeleteParent(DeleteParent event, Emitter<ParentState> emit) async {
    emit(ParentLoading());
    try {
      final response = await http.delete(
        Uri.parse('${Config.baseUrl}/parents/${event.id}'),
      );

      if (response.statusCode == 200) {
        emit(ParentSuccess('Parent deleted' as Parent));
      } else {
        emit(ParentError('Failed to delete parent'));
      }
    } catch (e) {
      emit(ParentError(e.toString()));
    }
  }
}
