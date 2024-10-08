import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
//import 'front/repositories/team_repository.dart';
import 'package:kermesse_land/utils/config.dart';
//import 'package:intl/intl.dart';
import 'package:jwt_decode/jwt_decode.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

//import '../../components/forms/participants_search_form.dart';
import '../../models/kermesse_model.dart';
//import '../../models/team_model.dart';
import '../../services/kermesse/kermesse_bloc.dart';
import '../../services/kermesse/kermesse_event.dart';
import '../../services/kermesse/kermesse_state.dart';
//import '../../services/team/team_bloc.dart';
//import '../../services/team/team_event.dart';
//import '../../services/team/team_state.dart';
import '../../models/user_model.dart';
//import '../team/team_manage_screen.dart';

/*class HackathonDetailPage extends StatefulWidget {
  final String id;
  final String token;

  const HackathonDetailPage({super.key, required this.id, required this.token});

  @override
  HackathonDetailPageState createState() => HackathonDetailPageState();
}

class HackathonDetailPageState extends State<HackathonDetailPage> {
  final Set<int> joinedTeams = <int>{};
  late int userId;
  late String username;
  User? currentUser;
  Hackathon? currentHackathon;
  late Future<void> fetchCurrentUserFuture;
  bool teamsInitialized = false;
  bool hasError = false;

  @override
  void initState() {
    super.initState();
    userId = _getUserIdFromToken(widget.token);
    username = _getUsernameFromToken(widget.token);
    fetchCurrentUserFuture = _fetchCurrentUser().then((_) {
      context.read<HackathonBloc>().add(FetchSingleHackathons(widget.token, widget.id));
    }).catchError((error) {
      setState(() {
        hasError = true;
      });
      if (kDebugMode) {
        print('Error fetching current user: $error');
      }
    });
  }

  int _getUserIdFromToken(String token) {
    Map<String, dynamic> decodedToken = Jwt.parseJwt(token);
    if (kDebugMode) {
      print('Decoded Token: $decodedToken');
    }
    return decodedToken['userId'];
  }

  String _getUsernameFromToken(String token) {
    Map<String, dynamic> decodedToken = Jwt.parseJwt(token);
    if (kDebugMode) {
      print('Decoded Token: $decodedToken');
    }
    return decodedToken['username'];
  }

  String formatDate(String dateString) {
    final date = DateFormat('yyyy-MM-dd').parse(dateString);
    return DateFormat('dd/MM/yyyy').format(date);
  }

  Future<void> _fetchCurrentUser() async {
    final url = '${Config.baseUrl}/user/me';
    final response = await http.get(
      Uri.parse(url),
      headers: {'Authorization': 'Bearer ${widget.token}'},
    );

    if (kDebugMode) {
      print('Fetching current user from $url');
      print('Response status: ${response.statusCode}');
      print('Response body: ${response.body}');
    }

    if (response.statusCode == 200) {
      setState(() {
        currentUser = User.fromJson(jsonDecode(response.body));
        if (kDebugMode) {
          print('Fetched current user: $currentUser');
        }
      });
    } else {
      throw Exception('Failed to load current user');
    }
  }

  /*void _initializeJoinedTeams(Hackathon hackathon) {
    if (teamsInitialized || currentUser == null) return; // Check if teams are already initialized and if currentUser is not null
    teamsInitialized = true; // Set initialized to true
    joinedTeams.clear();
    if (kDebugMode) {
      print('Initializing joined teams for user ID: $userId');
    }
    for (var team in hackathon.teams) {
      if (kDebugMode) {
        print('Checking team: ${team.name} with ID: ${team.id}');
      }
      if (team.users != null && team.users!.isNotEmpty) {
        for (var user in team.users!) {
          if (kDebugMode) {
            print(
                'Checking user: ${user.username} with ID: ${user.id} in team: ${team.name}');
          }
          if (user.id == userId) {
            joinedTeams.add(team.id);
            if (kDebugMode) {
              print('User $userId is part of team ${team.id}');
            }
          }
        }
      } else {
        if (kDebugMode) {
          print('No users found in team: ${team.name}');
        }
      }
    }
    if (kDebugMode) {
      print('Joined teams after initialization: $joinedTeams');
    }
  }

  void _updateTeamMembers(int teamId, bool isJoining) {
    if (currentUser == null) return; // Ensure currentUser is initialized
    setState(() {
      final team = currentHackathon!.teams.firstWhere((t) => t.id == teamId);
      if (isJoining) {
        team.users!.add(currentUser!);
        joinedTeams.add(teamId);
      } else {
        team.users!.removeWhere((u) => u.id == userId);
        joinedTeams.remove(teamId);
      }
      if (kDebugMode) {
        print(
            'Updated team members for team ID: $teamId, isJoining: $isJoining');
      }
    });
  }*/

  void _navigateToManagePage(BuildContext context, Team team) async {
    if (currentHackathon != null) {
      if (team.users != null && team.users!.any((user) => user.id == userId)) {
        final stepId = team.stepId ??
            (currentHackathon!.steps.isNotEmpty
                ? currentHackathon!.steps[0].id
                : null);
        final evaluationId = team.evaluationId ?? 0;

        final navigator = Navigator.of(context);
        final hackathonBloc = context.read<HackathonBloc>();

        final result = await navigator.push(
          MaterialPageRoute(
            builder: (context) => BlocProvider(
              create: (context) => TeamBloc(TeamRepository()),
              child: TeamManagePage(
                team: team,
                token: widget.token,
                evaluationId: evaluationId,
                stepId: stepId ?? 0,
                userId: userId,
                username: username,
              ),
            ),
          ),
        );

        if (result == 'left' && mounted) {
          hackathonBloc.add(FetchSingleHackathons(widget.token, widget.id));
        }
      } else {
        final scaffoldMessenger = ScaffoldMessenger.of(context);
        scaffoldMessenger.showSnackBar(
          const SnackBar(content: Text('Vous devez être membre de cette équipe pour accéder à la page de gestion.')),
        );
      }
    } else {
      final scaffoldMessenger = ScaffoldMessenger.of(context);
      scaffoldMessenger.showSnackBar(
        const SnackBar(content: Text('Hackathon non trouvé.')),
      );
    }
  }

  String _formatDate(String dateString) {
    final date = DateFormat('yyyy-MM-dd').parse(dateString);
    return DateFormat('dd/MM/yyyy').format(date);
  }

  String _formatDateTime(DateTime date) {
    //final date = DateFormat('yyyy-MM-dd').parse(dateString);
    return DateFormat('dd/MM/yyyy H:m').format(date);
  }

  String _getHackathonStatus(Hackathon hackathon) {
    final now = DateTime.now();
    final startDate = DateTime.parse(hackathon.startDate);
    final endDate = DateTime.parse(hackathon.endDate);
    final imminentDate = startDate.subtract(const Duration(days: 3));

    if (now.isAfter(endDate)) {
      return 'Terminé';
    } else if (now.isAfter(imminentDate) && now.isBefore(startDate)) {
      return 'Imminent';
    } else if (now.isAfter(startDate) && now.isBefore(endDate)) {
      return 'En cours';
    } else {
      return 'À venir';
    }
  }

  Color _getHackathonStatusColor(String status) {
    switch (status) {
      case 'Terminé':
        return Colors.red;
      case 'Imminent':
        return Colors.orange;
      case 'En cours':
        return Colors.green;
      case 'À venir':
        return Colors.blue;
      default:
        return Colors.black54;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Détails du hackathon'),
        centerTitle: true,
        backgroundColor: Colors.lime,
        actions: [
          IconButton(
            icon: const Icon(Icons.search),
            onPressed: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => ParticipantSearchForm(
                    hackathonId: int.parse(widget.id),
                    token: widget.token,
                  ),
                ),
              );
            },
          ),
        ],
      ),
      body: MultiBlocProvider(
        providers: [
          BlocProvider.value(
            value: context.read<HackathonBloc>(),
          ),
          BlocProvider(
            create: (context) => TeamBloc(TeamRepository()),
          ),
        ],
        child: FutureBuilder<void>(
          future: fetchCurrentUserFuture,
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.waiting) {
              return const Center(child: CircularProgressIndicator());
            } else if (snapshot.hasError || hasError) {
              return Center(child: Text('Error: ${snapshot.error ?? "Unknown error"}'));
            } else {
              return MultiBlocListener(
                listeners: [
                  BlocListener<HackathonBloc, HackathonState>(
                    listener: (context, hackathonState) {
                      if (hackathonState is HackathonLoaded) {
                        final hackathon = hackathonState.hackathons[0];
                        currentHackathon = hackathon;
                        if (kDebugMode) {
                          print('Fetched hackathon: ${hackathon.name}');
                        }
                        _initializeJoinedTeams(hackathon);
                        setState(() {});
                        if (kDebugMode) {
                          print('HackathonLoaded state emitted');
                        }
                      }
                    },
                  ),
                  BlocListener<TeamBloc, TeamState>(
                    listener: (context, state) {
                      final scaffoldMessenger = ScaffoldMessenger.of(context); // Capture ScaffoldMessenger before async call

                      if (state is TeamJoined) {
                        scaffoldMessenger.showSnackBar(
                          SnackBar(content: Text(state.message)),
                        );
                        _updateTeamMembers(state.teamId, true);
                      } else if (state is TeamLeft) {
                        scaffoldMessenger.showSnackBar(
                          SnackBar(content: Text(state.message)),
                        );
                        _updateTeamMembers(state.teamId, false);
                      } else if (state is TeamError) {
                        scaffoldMessenger.showSnackBar(
                          SnackBar(content: Text(state.error)),
                        );
                      }
                    },
                  ),
                ],
                child: BlocBuilder<HackathonBloc, HackathonState>(
                  builder: (context, hackathonState) {
                    if (hackathonState is HackathonLoading) {
                      return const Center(child: CircularProgressIndicator());
                    } else if (hackathonState is HackathonError) {
                      return Center(child: Text(hackathonState.message));
                    } else if (hackathonState is HackathonLoaded) {
                      final hackathon = hackathonState.hackathons[0];
                      final status = _getHackathonStatus(hackathon);
                      final startDate = _formatDate(hackathon.startDate);
                      final endDate = _formatDate(hackathon.endDate);
                      final userHasJoinedAnyTeam = joinedTeams.isNotEmpty;

                      return SingleChildScrollView(
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Stack(
                              children: [
                                Image.asset(
                                  "assets/logo.png",
                                  width: double.infinity,
                                  height: 200,
                                  fit: BoxFit.cover,
                                ),
                                Positioned(
                                  top: 16,
                                  right: 16,
                                  child: Chip(
                                    padding: const EdgeInsets.all(8.0),
                                    backgroundColor:
                                        _getHackathonStatusColor(status),
                                    label: Text(
                                      status,
                                      style: const TextStyle(
                                        color: Colors.white,
                                        fontSize: 18,
                                      ),
                                    ),
                                    side: BorderSide.none,
                                  ),
                                ),
                              ],
                            ),
                            Padding(
                              padding: const EdgeInsets.all(16.0),
                              child: Column(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [

                                  Card(
                                    color: Colors.lime.withOpacity(0.8),
                                    shape: RoundedRectangleBorder(
                                      borderRadius: BorderRadius.circular(15.0),
                                    ),
                                    elevation: 5,
                                    margin: const EdgeInsets.all(10),
                                    child: Padding(
                                      padding: const EdgeInsets.all(16.0),
                                      child: Column(
                                        crossAxisAlignment: CrossAxisAlignment.start,
                                        children: [
                                          Text(
                                            hackathon.name,
                                            style: const TextStyle(
                                              fontSize: 24,
                                              fontWeight: FontWeight.bold,
                                              color: Colors.white
                                            ),
                                          ),
                                          const SizedBox(height: 8),
                                          Row(
                                            children: [
                                              const SizedBox(height: 8),
                                              const Icon(Icons.location_pin, color: Colors.white),
                                              const SizedBox(width: 8),
                                              Flexible(
                                                child: Text(
                                                  hackathon.location,
                                                  style: const TextStyle(
                                                    fontSize: 24,
                                                    fontWeight: FontWeight.bold,
                                                  ),
                                                  overflow: TextOverflow.ellipsis,
                                                ),
                                              ),
                                            ],
                                          ),
                                          const SizedBox(height: 8),
                                          Row(
                                            children: [
                                              const Icon(Icons.date_range, color: Colors.white),
                                              const SizedBox(width: 8),
                                              Text(
                                                'Du $startDate au $endDate',
                                                style: const TextStyle(
                                                  fontSize: 16,
                                                  color: Colors.white,
                                                ),
                                              ),
                                            ],
                                          ),
                                        ],
                                      ),
                                    ),
                                  ),

                                  const SizedBox(height: 16),
                                  SizedBox(
                                    width: double.infinity, // Permet à la Card d'occuper toute la largeur disponible
                                    child: Card(
                                      shape: RoundedRectangleBorder(
                                        borderRadius: BorderRadius.circular(15.0),
                                      ),
                                      elevation: 5,
                                      margin: const EdgeInsets.all(10),
                                      child: Padding(
                                        padding: const EdgeInsets.all(16.0),
                                        child: Column(
                                          crossAxisAlignment: CrossAxisAlignment.start,
                                          children: [
                                            const Text(
                                              'Description: ',
                                              style: TextStyle(
                                                fontSize: 20,
                                                fontWeight: FontWeight.bold,
                                              ),
                                            ),
                                            const SizedBox(height: 16),
                                            Text(
                                              hackathon.description,
                                              style: const TextStyle(
                                                fontSize: 16,
                                              ),
                                            ),
                                            const SizedBox(height: 24),
                                            const Text(
                                              'Les équipes: ',
                                              style: TextStyle(
                                                fontSize: 20,
                                                fontWeight: FontWeight.bold,
                                              ),
                                            ),
                                            const SizedBox(height: 8),
                                            Wrap(
                                              spacing: 8.0,
                                              runSpacing: 4.0,
                                              children: hackathon.teams
                                                  .map(
                                                    (participant) => Chip(
                                                  label: Text(participant.name),
                                                ),
                                              )
                                                  .toList(),
                                            ),
                                            const SizedBox(height: 16),
                                            Column(
                                              children: hackathon.teams.map((team) {
                                                final isUserInTeam = joinedTeams.contains(team.id);
                                                return ListTile(
                                                  title: Text(team.name),
                                                  subtitle: Text(
                                                    team.users != null && team.users!.isNotEmpty
                                                        ? 'Membres: ${team.users!.map((user) => user.username).join(', ')}'
                                                        : 'Aucun membre',
                                                  ),
                                                  trailing: isUserInTeam ?
                                                  ElevatedButton(
                                                    onPressed: () {
                                                      context.read<TeamBloc>().add(LeaveTeam(team.id, widget.token));
                                                    },
                                                    style: ElevatedButton.styleFrom(
                                                      backgroundColor: isUserInTeam ? Colors.red : Colors.lime,
                                                    ),
                                                    child: const Text(
                                                      'Quitter',
                                                      style: TextStyle(color: Colors.white),
                                                    ),
                                                  ) : userHasJoinedAnyTeam
                                                    ? null
                                                    : ElevatedButton(
                                                    onPressed: () {
                                                      context.read<TeamBloc>().add(JoinTeam(team.id, widget.token));
                                                    },
                                                    style: ElevatedButton.styleFrom(
                                                      backgroundColor: isUserInTeam ? Colors.red : Colors.lime,
                                                    ),
                                                    child: const Text(
                                                      'Rejoindre',
                                                      style: TextStyle(color: Colors.white),
                                                    ),
                                                  ),
                                                  onTap: () => _navigateToManagePage(context, team),
                                                );
                                              }).toList(),
                                            ),
                                          ],
                                        ),
                                      ),
                                    ),
                                  ),


                                  const SizedBox(height: 8,),
                                  if (hackathon.steps.isNotEmpty) ...[
                                    Column(
                                      children: [
                                        const Text('Steps:', style: TextStyle(fontSize: 18)),
                                        ListView.builder(
                                          shrinkWrap: true, // Add this line
                                          physics: const NeverScrollableScrollPhysics(), // Add this line
                                          itemCount: hackathon.steps.length,
                                          itemBuilder: (context, index) {
                                            final step = hackathon.steps[index];
                                            return ListTile(
                                              title: Text(step.title),
                                              subtitle: Text('Position: ${step.position}, Deadline: ${_formatDateTime(step.deadLineDate)}'),
                                            );
                                          },
                                        ),
                                      ],
                                    ),
                                  ] else ...[
                                    const Center(child: Text('Aucune étape disponible.', style: TextStyle(fontSize: 18))),
                                  ],
                                ],
                              ),
                            ),
                          ],
                        ),
                      );
                    } else {
                      return const Center(
                          child: Text('Aucune information disponible.'));
                    }
                  },
                ),
              );
            }
          },
        ),
      ),
    );
  }
}*/

