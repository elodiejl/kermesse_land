import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
//import 'package:front/screens/kermesse/kermesse_teams_screen.dart';
//import 'package:intl/intl.dart';
//import '../../components/forms/add_hackathon_form.dart';
import '../../services/kermesse/kermesse_bloc.dart';
import '../../services/kermesse/kermesse_event.dart';
import '../../services/kermesse/kermesse_state.dart';

class KermesseScreen extends StatefulWidget {
  final String token;
  //final String googleApiKey = dotenv.env['GOOGLE_PLACES_API_KEY'] ?? 'YOUR_FALLBACK_API_KEY';

  const KermesseScreen({super.key, required this.token});

  @override
  KermesseScreenState createState() => KermesseScreenState();
}

class KermesseScreenState extends State<KermesseScreen> {
  late KermesseBloc _kermesseBloc;

  @override
  void initState() {
    super.initState();
    _kermesseBloc = KermesseBloc();
    _kermesseBloc.add(FetchKermesseForUser(widget.token));
  }

  void _showAddKermesseDialog(BuildContext context, String token) {
    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: const Text('Ajouter une Kermesse'),
          content: SizedBox(
            width: MediaQuery.of(context).size.width * 0.85,
            /*child: AddKermesseForm(
              token: token,
              googleApiKey: widget.googleApiKey,
            ),*/
          ),
          backgroundColor: Colors.white,
        );
      },
    ).then((_) {
      _kermesseBloc.add(FetchKermesseForUser(token)); // Refresh kermesses after the dialog is closed
    });
  }

  String formatDate(String dateString) {
    final date = DateTime.parse(dateString);
    return '${date.day}/${date.month}/${date.year}';
  }

  Map<String, dynamic> getStatus(String dateString, String dateFinaleString) {
    final kermesseDate = DateTime.parse(dateString);
    //final kermesseDateFinale = DateTime.parse(dateFinaleString);
    final currentDate = DateTime.now();
    final difference = kermesseDate.difference(currentDate).inDays;
    //final differencedf = kermesseDateFinale.difference(currentDate).inDays;

    if (difference == 0) {
      return {'text': 'En cours', 'color': Colors.orangeAccent};
    } else if (difference < 0) {
      return {'text': 'Terminé', 'color': Colors.red};
    } else if (difference < 3) {
      return {'text': 'Imminent', 'color': Colors.orange};
    } else {
      return {'text': 'À venir', 'color': Colors.green};
    }
  }

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => _kermesseBloc,
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Mes Kermesses'),
          backgroundColor: Colors.cyan,
        ),
        body: BlocListener<KermesseBloc, KermesseState>(
          listener: (context, state) {
            if (state is KermesseLoaded) {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(content: Text('Kermesse mise à jour')),
              );
            }
          },
          child: BlocBuilder<KermesseBloc, KermesseState>(
            builder: (context, state) {
              if (state is KermesseLoading) {
                return const Center(child: CircularProgressIndicator());
              } else if (state is KermesseLoaded) {
                if (kDebugMode) {
                  print(state.kermesses.length);
                }
                if (state.kermesses.isEmpty) {
                  return const Center(child: Text('Aucune kermesse trouvée.'));
                }
                return ListView.builder(
                  itemCount: state.kermesses.length,
                  itemBuilder: (context, index) {
                    final kermesse = state.kermesses[index];
                    return Padding(
                      padding: const EdgeInsets.all(8.0),
                      child: Card(
                        elevation: 4.0,
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(10.0),
                        ),
                        color: Colors.white,
                        child: InkWell(
                          onTap: () {
                            // Action quand on clique sur une kermesse
                          },
                          child: Padding(
                            padding: const EdgeInsets.all(10.0),
                            child: Row(
                              children: [
                                const SizedBox(width: 10),
                                Expanded(
                                  child: Column(
                                    crossAxisAlignment: CrossAxisAlignment.start,
                                    children: [
                                      Row(
                                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                                        children: [
                                          Text(
                                            kermesse.name,
                                            style: const TextStyle(
                                              fontSize: 18.0,
                                              fontWeight: FontWeight.bold,
                                            ),
                                          ),
                                          Container(
                                            padding: const EdgeInsets.symmetric(
                                              horizontal: 8.0,
                                              vertical: 4.0,
                                            ),
                                            decoration: BoxDecoration(
                                              color: getStatus(kermesse.date, kermesse.date)['color'],
                                              borderRadius: BorderRadius.circular(12.0),
                                            ),
                                            child: Text(
                                              getStatus(kermesse.date, kermesse.date)['text'],
                                              style: const TextStyle(
                                                color: Colors.white,
                                                fontSize: 12.0,
                                                fontWeight: FontWeight.bold,
                                              ),
                                            ),
                                          ),
                                        ],
                                      ),
                                      const SizedBox(height: 5),
                                      Text(
                                        '${formatDate(kermesse.date)} - ${kermesse.location}',
                                        style: TextStyle(
                                          color: Colors.grey[600],
                                        ),
                                      ),
                                      const SizedBox(height: 5),
                                      Text(
                                        kermesse.location,
                                        maxLines: 2,
                                        overflow: TextOverflow.ellipsis,
                                        style: TextStyle(
                                          color: Colors.grey[800],
                                        ),
                                      ),
                                    ],
                                  ),
                                ),
                              ],
                            ),
                          ),
                        ),
                      ),
                    );
                  },
                );
              } else if (state is KermesseError) {
                return Center(child: Text('Erreur: ${state.message}'));
              } else {
                return const Center(child: Text('Aucune donnée disponible'));
              }
            },
          ),
        ),
        floatingActionButton: FloatingActionButton(
          backgroundColor: Colors.cyan,
          hoverColor: Colors.cyan[700],
          onPressed: () {
            //_showAddKermesseDialog(context);
          },
          tooltip: 'Ajouter une kermesse',
          child: const Icon(Icons.add),
        ),
      ),
    );
  }
}