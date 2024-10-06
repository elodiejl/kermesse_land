import 'dart:convert';
import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;
import 'package:http_parser/http_parser.dart';

import '../models/user_model.dart';
import '../utils/config.dart';

class AuthenticationService {
  late String _token;
  String get token => _token;

  Future<bool> login(String email, String password) async {
    try {
      final response = await http.post(
        Uri.parse('${Config.baseUrl}/user/login'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, String>{
          'email': email,
          'password': password,
        }),
      );

      if (response.statusCode == 200) {
        if (kDebugMode) {
          print('Login successful, token: ${jsonDecode(response.body)['token']}');
        }
        _token = jsonDecode(response.body)['token'];
        return true;
      } else {
        if (kDebugMode) {
          print('Login failed: ${response.body}');
        }
        return false;
      }
    } catch (e) {
      if (kDebugMode) {
        print('Login error: $e');
      }
      return false;
    }
  }

  Future<void> register({
    required String username,
    required String lastName,
    required String firstName,
    required String email,
    required String password,
    required String role
  }) async {
    var uri = Uri.parse('${Config.baseUrl}/user/register');
    var request = http.MultipartRequest('POST', uri)
      ..fields['username'] = username
      ..fields['last_name'] = lastName
      ..fields['first_name'] = firstName
      ..fields['email'] = email
      ..fields['password'] = password
      ..fields['role'] = role;



    var response = await request.send();
    if (response.statusCode == 201) {
      if (kDebugMode) {
        print('Registration successful');
      }
    } else {
      if (kDebugMode) {
        print('Registration failed: ${response.reasonPhrase}');
      }
      throw Exception('Failed to register');
    }
  }

  // Destroy the token
  Future<bool> logout() async {
    _token = '';
    return true;
  }

  Future<User?> getUserDetails(String token) async {
    try {
      final response = await http.get(
        Uri.parse('${Config.baseUrl}/user/me'),
        headers: <String, String>{
          'Authorization': 'Bearer $token',
        },
      );

      if (kDebugMode) {
        print('Get user details response status: ${response.statusCode}');
        print('Get user details response body: ${response.body}');
      }

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        return User.fromJson(responseBody);
      } else {
        if (kDebugMode) {
          print('Failed to get user details: ${response.body}');
        }
        return null;
      }
    } catch (e) {
      if (kDebugMode) {
        print('Error fetching user details: $e');
      }
      return null;
    }
  }
}
