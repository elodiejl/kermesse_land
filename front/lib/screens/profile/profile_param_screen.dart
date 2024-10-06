import 'dart:io';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../services/authentication/authentication_bloc.dart';
import '../../services/authentication_service.dart';
import '../../services/user/user_bloc.dart';
import '../../services/user/user_event.dart';
import '../../services/user/user_state.dart';
import '../../models/user_model.dart';


class ProfilParamPage extends StatefulWidget {
  final String token;

  const ProfilParamPage({super.key, required this.token});

  @override
  ProfilParamPageState createState() => ProfilParamPageState();
}

class ProfilParamPageState extends State<ProfilParamPage> {
  late TextEditingController _firstNameController;
  late TextEditingController _lastNameController;
  late TextEditingController _usernameController;
  late TextEditingController _emailController;
  late String roleText;

  @override
  void initState() {
    super.initState();
    _firstNameController = TextEditingController();
    _lastNameController = TextEditingController();
    _usernameController = TextEditingController();
    _emailController = TextEditingController();
    roleText = '';
  }

  @override
  void dispose() {
    _firstNameController.dispose();
    _lastNameController.dispose();
    _usernameController.dispose();
    _emailController.dispose();
    roleText = '';
    super.dispose();
  }


  Future<void> _saveChanges(BuildContext context) async {
    final userBloc = BlocProvider.of<UserBloc>(context);
    final user = userBloc.state is UserLoaded ? (userBloc.state as UserLoaded).user : null;

    if (user != null) {

      final updatedUser = UserUpdate(
        id: user.id,
        firstName: _firstNameController.text,
        lastName: _lastNameController.text,
        username: _usernameController.text,
        email: _emailController.text,
      );

      userBloc.add(UpdateUser(widget.token, updatedUser));

      // Listen to the UserBloc state changes
      userBloc.stream.listen((state) {
        if (state is UserLoaded) {
          // Show a snackbar when the user is updated successfully
          ScaffoldMessenger.of(context)
            ..removeCurrentSnackBar()
            ..showSnackBar(
              const SnackBar(content: Text("User updated successfully")),
            );
        } else if (state is UserError) {
          // Show a snackbar if there's an error updating the user
          ScaffoldMessenger.of(context)
            ..removeCurrentSnackBar()
            ..showSnackBar(
              SnackBar(content: Text("Error updating user: ${state.message}")),
            );
        }
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider<UserBloc>(
          create: (context) => UserBloc()..add(FetchUser(widget.token)),
        ),
        BlocProvider<AuthenticationBloc>(
          create: (context) => AuthenticationBloc(AuthenticationService()),
        ),
      ],
      child: Scaffold(
        appBar: AppBar(
          title: const Text("Profil", textAlign: TextAlign.center),
          centerTitle: true,
          backgroundColor: Colors.cyan,
        ),
        body: BlocBuilder<UserBloc, UserState>(
          builder: (context, state) {
            if (state is UserLoading) {
              return const Center(child: CircularProgressIndicator());
            } else if (state is UserLoaded) {
              _firstNameController.text = state.user.firstName;
              _lastNameController.text = state.user.lastName;
              _usernameController.text = state.user.username;
              _emailController.text = state.user.email;

              /*String roleText;

              switch (state.user.roles) {
                case 4:
                  roleText = "Élève";
                  break;
                case 8:
                  roleText = "Parent";
                  break;
                case 1:
                  roleText = "Organisateur";
                  break;
                case 2:
                  roleText = "Admin";
                  break;
                case 16:
                  roleText = "Gérant de stand";
                  break;
                default:
                  roleText = "Rôle Inconnu";
              }*/

              return SingleChildScrollView(
                child: Center(
                  child: Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: Column(
                      children: [
                        const SizedBox(height: 20),
                        const CircleAvatar(
                          radius: 60,
                          backgroundImage: AssetImage('assets/user-icon.png'),
                        ),
                        const Padding(
                          padding: EdgeInsets.all(8.0),
                          child: Divider(
                            color: Colors.blueGrey,
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.all(8),
                          child: TextFormField(
                            controller: _firstNameController,
                            decoration: const InputDecoration(
                              labelText: 'Firstname',
                            ),
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.all(8),
                          child: TextFormField(
                            controller: _lastNameController,
                            decoration: const InputDecoration(
                              labelText: 'Lastname',
                            ),
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.all(8),
                          child: TextFormField(
                            controller: _usernameController,
                            decoration: const InputDecoration(
                              labelText: 'Username',
                            ),
                            readOnly: true,
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.all(8),
                          child: TextFormField(
                            controller: _emailController,
                            decoration: const InputDecoration(
                              labelText: 'Email',
                            ),
                          ),
                        ),
                        const Padding(
                          padding: EdgeInsets.all(8),
                          child: Row(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              SizedBox(
                                width: 80,
                                child: Text(
                                  'Rôle: ',
                                  style: TextStyle(
                                    color: Colors.black,
                                    fontWeight: FontWeight.normal,
                                  ),
                                ),
                              ),
                              //Text(roleText, style: TextStyle(fontWeight: FontWeight.bold)),
                            ],
                          ),
                        ),
                        const SizedBox(height: 200),
                        Padding(
                          padding: const EdgeInsets.only(bottom: 16.0),
                          child: Align(
                            alignment: Alignment.bottomCenter,
                            child: ElevatedButton(
                              onPressed: () => _saveChanges(context),
                              style: ElevatedButton.styleFrom(
                                backgroundColor: Colors.cyan,
                                padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 15),
                              ),
                              child: const Text(
                                "Save changes",
                                style: TextStyle(fontSize: 18, color: Colors.white),
                              ),
                            ),
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              );
            } else if (state is UserError) {
              return Center(child: Text("Error: ${state.message}"));
            } else {
              return const Center(
                  child: Text("An unknown error occurred."));
            }
          },
        ),
      ),
    );
  }
}