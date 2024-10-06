import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
//import 'package:image_picker/image_picker.dart';
import 'dart:io';

import '../../services/register/register_bloc.dart';
import '../../services/register/register_event.dart';
import '../../services/register/register_state.dart';
import '../../widgets/buttons/button_base.dart';

class RegisterPage extends StatefulWidget {
  const RegisterPage({super.key});

  @override
  RegisterScreenState createState() => RegisterScreenState();
}

class RegisterScreenState extends State<RegisterPage> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _lastNameController = TextEditingController();
  final TextEditingController _firstNameController = TextEditingController();
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  String? _selectedRole;

  void _onRegisterButtonPressed() {
    if (_formKey.currentState!.validate()) {
      BlocProvider.of<RegistrationBloc>(context).add(
        SignUpButtonPressed(
          username: _usernameController.text,
          lastName: _lastNameController.text,
          firstName: _firstNameController.text,
          email: _emailController.text,
          password: _passwordController.text,
          role: _selectedRole ?? '',
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Padding(
          padding: const EdgeInsets.all(90.0),
          child: Image.asset(
            'assets/logo.png',
            height: 100.0,
          ),
        ),
      ),
      body: BlocListener<RegistrationBloc, RegistrationState>(
        listener: (context, state) {
          if (state is RegistrationSuccess) {
            ScaffoldMessenger.of(context)
              ..removeCurrentSnackBar()
              ..showSnackBar(const SnackBar(content: Text('Inscription réussie! Vous pouvez vous connecter maintenant.')));
          }
          if (state is RegistrationFailure) {
            ScaffoldMessenger.of(context)
              ..removeCurrentSnackBar()
              ..showSnackBar(SnackBar(content: Text('Erreur d\'inscription: ${state.error}')));
          }
        },
        child: Form(
          key: _formKey,
          child: SingleChildScrollView(
            padding: const EdgeInsets.all(20),
            child: Column(
              children: [
                const SizedBox(height: 30),
                const Text(
                  'Bienvenue sur Kermesse Land',
                  style: TextStyle(
                    fontSize: 20.0,
                    fontWeight: FontWeight.w500,
                  ),
                ),
                const SizedBox(height: 10),
                const Center(
                  child: Text(
                    'Créer un compte',
                    style: TextStyle(
                      fontSize: 28.0,
                      fontWeight: FontWeight.bold,
                    ),
                    textAlign: TextAlign.center,
                  ),
                ),
                const SizedBox(height: 30),
                const CircleAvatar(
                  radius: 60,
                  backgroundImage: AssetImage('assets/user-icon.png'),
                ),
                TextFormField(
                  controller: _usernameController,
                  decoration: const InputDecoration(labelText: 'Nom d\'utilisateur'),
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Veuillez entrer un nom d\'utilisateur';
                    }
                    return null;
                  },
                ),
                TextFormField(
                  controller: _lastNameController,
                  decoration: const InputDecoration(labelText: 'Nom'),
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Veuillez entrer un nom';
                    }
                    return null;
                  },
                ),
                TextFormField(
                  controller: _firstNameController,
                  decoration: const InputDecoration(labelText: 'Prénom'),
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Veuillez entrer un prénom';
                    }
                    return null;
                  },
                ),
                TextFormField(
                  controller: _emailController,
                  decoration: const InputDecoration(labelText: 'Email'),
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Veuillez entrer un email';
                    }
                    return null;
                  },
                ),
                TextFormField(
                  controller: _passwordController,
                  decoration: const InputDecoration(labelText: 'Mot de passe'),
                  obscureText: true,
                  validator: (value) {
                    if (value!.isEmpty) {
                      return 'Veuillez entrer un mot de passe';
                    }
                    return null;
                  },
                ),
                const SizedBox(height: 10),
                DropdownButtonFormField<String>(
                  value: _selectedRole,
                  decoration: const InputDecoration(
                    labelText: 'Choisissez votre rôle',
                  ),
                  items: const <DropdownMenuItem<String>>[
                    DropdownMenuItem<String>(value: "4", child: Text('Élève')),
                    DropdownMenuItem<String>(value: "8", child: Text('Parent')),
                    DropdownMenuItem<String>(value: "1", child: Text('Organisateur')),
                    DropdownMenuItem<String>(value: "16", child: Text('Gérant de stand')),
                  ],
                  onChanged: (String? newValue) {
                    setState(() {
                      _selectedRole = newValue;
                    });
                  },
                  validator: (value) {
                    if (value == null || value.isEmpty) {
                      return 'Veuillez sélectionner un rôle';
                    }
                    return null;
                  },
                ),

                const SizedBox(height: 30),
                SizedBox(
                  width: MediaQuery.of(context).size.width,
                  child: ButtonBase(
                    text: 'S\'inscrire',
                    onPressed: _onRegisterButtonPressed,
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
