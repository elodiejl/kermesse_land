import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../services/stand/stand_bloc.dart';
import '../../services/stand/stand_event.dart';
import '../../services/stand/stand_state.dart';
import 'package:kermesse_land/models/stand_model.dart';

class StandsByKermesseScreen extends StatefulWidget {
  final String token;
  final int kermesseId;

  const StandsByKermesseScreen({super.key, required this.token, required this.kermesseId});

  @override
  StandsByKermesseScreenState createState() => StandsByKermesseScreenState();
}

class StandsByKermesseScreenState extends State<StandsByKermesseScreen> {
  late StandBloc _standsBloc;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    _standsBloc = BlocProvider.of<StandBloc>(context);
    _standsBloc.add(FetchStandsByKermesse(widget.token, widget.kermesseId)); // Événement pour charger les stands
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Stands de la Kermesse'),
        backgroundColor: Colors.cyan,
      ),
      body: BlocBuilder<StandBloc, StandState>(
        builder: (context, state) {
          if (state is StandLoading) {
            return const Center(child: CircularProgressIndicator());
          } else if (state is StandError) {
            return Center(child: Text(state.message));
          } else if (state is StandLoaded) {
            return ListView.builder(
              itemCount: state.stands.length,
              itemBuilder: (context, index) {
                final stand = state.stands[index];

                return Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: Card(
                    elevation: 4.0,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10.0),
                    ),
                    child: ListTile(
                      title: Text(stand.name),
                      subtitle: Text('Coût de participation: ${stand.participationCost} jetons'), // Détails supplémentaires
                      onTap: () {
                        // Action quand on clique sur le stand
                        // Naviguer vers une nouvelle page ou afficher des détails supplémentaires
                      },
                    ),
                  ),
                );
              },
            );
          }
          return Container();
        },
      ),
    );
  }
}
