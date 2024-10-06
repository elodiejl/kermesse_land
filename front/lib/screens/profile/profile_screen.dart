import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/screens/profile/profile_param_screen.dart';
import '../../services/logout/logout_bloc.dart';
import '../../services/logout/logout_event.dart';
import '../../services/logout/logout_state.dart';
import '../../services/user/user_bloc.dart';
import '../../services/user/user_event.dart';
import '../../services/user/user_state.dart';
import '../../services/authentication_service.dart';
import '../token_purchase/token_management_screen.dart';

class ProfileScreen extends StatelessWidget {
  final String token;

  const ProfileScreen({super.key, required this.token});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider<UserBloc>(
          create: (context) => UserBloc()..add(FetchUser(token)),
        ),
        BlocProvider<LogoutBloc>(
          create: (context) => LogoutBloc(AuthenticationService()), // Correctly refer to AuthenticationService
        ),
      ],
      child: BlocListener<LogoutBloc, LogoutState>(
        listener: (context, state) {
          if (state is LogoutSuccess) {
            // Navigate to login page on successful logout
            Navigator.pushReplacementNamed(context, '/login');
          } else if (state is LogoutFailure) {
            // Handle logout failure here if needed
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text('Failed to logout: ${state.error}')),
            );
          }
        },
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
                String roleText;

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
                  default:
                    roleText = "Rôle Inconnu"; // Valeur par défaut
                }
                return SingleChildScrollView(
                  child: Center(
                    child: Padding(
                      padding: const EdgeInsets.all(16.0),
                      child: Column(
                        mainAxisAlignment: MainAxisAlignment.start,
                        crossAxisAlignment: CrossAxisAlignment.center,
                        children: [
                          CircleAvatar(
                            radius: 60,
                            backgroundImage: const AssetImage('assets/user-icon.png'),
                            onBackgroundImageError: (exception, stackTrace) {
                              // Gestion des erreurs de chargement d'image ici
                              if (kDebugMode) {
                                print('Image loading error: $exception');
                              }
                              // Afficher une image de remplacement ou un message d'erreur
                            },
                          ),
                          const SizedBox(height: 20),
                          Text(
                            roleText,
                            style: Theme.of(context).textTheme.titleLarge?.copyWith(
                              color: Colors.blueGrey[800],
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                          const SizedBox(height: 10),
                          Text(
                            "${state.user.firstName} ${state.user.lastName}",
                            style: Theme.of(context).textTheme.titleLarge?.copyWith(
                              color: Colors.blueGrey[800],
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                          const SizedBox(height: 10),
                          Text(
                            state.user.email,
                            style: Theme.of(context).textTheme.titleMedium?.copyWith(
                              color: Colors.blueGrey[600],
                            ),
                          ),
                          const SizedBox(height: 30),
                          if (state.user.roles == 8) // 8 pour Parent
                            buildProfileItem(
                              context,
                              icon: Icons.generating_tokens_outlined,
                              title: 'Gestion des tokens',
                              subtitle: 'Pour acheter et accéder à votre historique d\'achat',
                              onTap: () {
                                Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) => TokenManagementPage(
                                        token: token
                                    ),
                                  ),
                                );
                              },
                            ),
                          buildProfileItem(
                            context,
                            icon: Icons.settings,
                            title: 'Paramètres',
                            subtitle: 'Modifier vos paramètres de compte',
                            onTap: () {
                              Navigator.push(
                                context,
                                MaterialPageRoute(
                                  builder: (context) => ProfilParamPage(
                                    token: token,
                                  ),
                                ),
                              );
                            },
                          ),
                          /*if (state.user.roles == 4)
                            buildProfileItem(
                              context,
                              icon: Icons.admin_panel_settings,
                              title: 'Admin',
                              subtitle: 'Accéder aux fonctionnalités administrateur',
                              onTap: () {
                                // Navigate to admin settings page
                              },
                            ),*/
                          const SizedBox(height: 30),
                          ElevatedButton(
                            onPressed: () {
                              context.read<LogoutBloc>().add(LogoutRequested());
                            },
                            style: ElevatedButton.styleFrom(
                              backgroundColor: Colors.red,
                              padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 15),
                            ),
                            child: const Text(
                              "Déconnexion",
                              style: TextStyle(fontSize: 18, color: Colors.white),
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                );
              } else if (state is UserError) {
                return Center(child: Text("Erreur : ${state.message}"));
              } else {
                return const Center(child: Text("Une erreur inconnue s'est produite."));
              }
            },
          ),
        ),
      ),
    );
  }

  Widget buildProfileItem(BuildContext context,
      {required IconData icon,
        required String title,
        required String subtitle,
        required VoidCallback onTap}) {
    return Card(
      margin: const EdgeInsets.symmetric(vertical: 10),
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
      elevation: 5,
      child: ListTile(
        leading: Icon(icon, color: Colors.blueGrey[700]),
        title: Text(
          title,
          style: Theme.of(context).textTheme.titleMedium?.copyWith(color: Colors.blueGrey[800]),
        ),
        subtitle: Text(
          subtitle,
          style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: Colors.blueGrey[600]),
        ),
        trailing: const Icon(Icons.arrow_forward_ios, color: Colors.blueGrey),
        onTap: onTap,
      ),
    );
  }
}
