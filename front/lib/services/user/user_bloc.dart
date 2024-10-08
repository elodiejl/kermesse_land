import 'dart:convert';

import 'package:bloc/bloc.dart';
import 'package:kermesse_land/services/user/user_event.dart';
import 'package:kermesse_land/services/user/user_state.dart';
import 'package:http/http.dart' as http;

import '../../models/user_model.dart';
import '../../utils/config.dart';

class UserBloc extends Bloc<UserEvent, UserState> {
  UserBloc() : super(UserInitial()) {
    on<FetchUser>(_onFetchUser);
    on<UpdateUser>(_onUpdateUser);
  }

  void _onFetchUser(FetchUser event, Emitter<UserState> emit) async {
    emit(UserLoading());
    try {
      final user = await _fetchUser(event.token);
      emit(UserLoaded(user));
    } catch (e) {
      emit(UserError(e.toString()));
    }
  }

  Future<User> _fetchUser(String token) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/user/me'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      return User.fromJson(jsonDecode(response.body));
    } else {
      throw Exception('Failed to load user');
    }
  }
  void _onUpdateUser(UpdateUser event, Emitter<UserState> emit) async {
    emit(UserLoading());
    try {
      final user = await _updateUser(event.token, event.updatedUser);
      emit(UserLoaded(user));
    } catch (e) {
      emit(UserError(e.toString()));
    }
  }

  Future<User> _updateUser(String token, UserUpdate updatedUser) async {
    var uri = Uri.parse('${Config.baseUrl}/user/me');
    var request = http.MultipartRequest('PUT', uri);
    request.headers['Authorization'] = 'Bearer $token';

    // Add text fields
    request.fields['first_name'] = updatedUser.firstName ?? '';
    request.fields['last_name'] = updatedUser.lastName ?? '';
    request.fields['email'] = updatedUser.email ?? '';

    var response = await request.send();

    if (response.statusCode == 200) {
      var responseString = await response.stream.bytesToString();
      return User.fromJson(json.decode(responseString));
    } else {
      var responseString = await response.stream.bytesToString();
      var errorResponse = json.decode(responseString);
      var errorMessage = errorResponse['error'] ?? 'Failed to update user';
      throw Exception(errorMessage);
    }
  }

}
