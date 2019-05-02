package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func NewHttpStub(response string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, response)
	}))
	return server
}

const errorHtmlOutput = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
	"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="fr-FR" lang="fr-FR">
					<head>
		                            				
                	<title>TCL - Service Attestation d abonnement</title>
    
    
    
    				<meta name="keywords" content="Service Attestation d abonnement" />
			    	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
            <meta name="Content-Type" content="text/html; charset=utf-8" />

            <meta name="Content-language" content="fr-FR" />

                    <meta name="author" content="Micropole" />
    
                <meta name="copyright" content="Tcl.fr" />
    
                <meta name="description" content="TCL, réseau des Transports en Commun Lyonnais. Plan des lignes de métro, bus, tramway et funiculaire sur Lyon, Villeurbanne et l'agglomération lyonnaise ; horaires des bus, trams et métros ; trajets sur-mesure : calculer un itinéraire ; tarif des tickets de bus, carnets, abonnements, cartes Técély" />
    
                <meta name="keywords" content="TCL, transports en commun lyonnais, bus rhône, tram lyon, métro ville lyon, funiculaire lyon, tramway lyon, plan lignes, lignes métro, lignes bus, lignes tramways, lignes funiculaire, transports en commun lyon, transports en commun Villeurbanne, transports en commun Rhône, horaires bus, horaires trams, horaires métros, calcul trajets, calcul itinéraire, tarif tickets bus, tarif abonnement, tarif carte abonnement, tarif carte Técély, bons plans lyon" />
    
        <meta name="MSSmartTagsPreventParsing" content="TRUE" />
		<link rel="stylesheet" type="text/css" href="/var/tcl/cache/public/stylesheets/6f4a838f23e9676c03c193a31c890c93_1554111003_screen.css" media="screen" />


<link rel="stylesheet" type="text/css" href="/var/tcl/cache/public/stylesheets/96a690e8a6d5a55c1f81991515a26111_1554111003_print.css" media="print" />



        <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,400,700,800" rel="stylesheet">

<!--[if IE]>
		<link href="/extension/tcl/design/fr/stylesheets/ie-tous.css" rel="stylesheet" type="text/css" media="screen" />
		<link href="/extension/tcl/design/fr/stylesheets/ie-tous_print.css" rel="stylesheet" type="text/css" media="print" />
<![endif]-->

<!--[if lt IE 7]>
		<link href="/extension/tcl/design/fr/stylesheets/ie6.css" rel="stylesheet" type="text/css" media="screen" />
<![endif]-->

<!--[if IE 7]>
		<link href="/extension/tcl/design/fr/stylesheets/ie7.css" rel="stylesheet" type="text/css" media="screen" />
<![endif]-->
<!--[if IE 8]>
		<link href="/extension/tcl/design/fr/stylesheets/ie8.css" rel="stylesheet" type="text/css" media="screen" />
<![endif]-->
<!--[if gt IE 7]>
		<link href="/extension/tcl/design/fr/stylesheets/ie-tous.css" rel="stylesheet" type="text/css" media="screen" />
<![endif]-->
		<script type="text/javascript" src="/var/tcl/cache/public/javascript/06808eb35470cc7c2a9fc0609802e2e6_1554111002.js" charset="utf-8"></script>


<script type="text/javascript" src="/var/tcl/cache/public/javascript/37cff028e4be8809d6322fe585f6f543_1554111003.js" charset="utf-8"></script>


		
				
		<link rel="Shortcut icon" href="/extension/tcl/design/fr/images/favicon.ico" type="image/x-icon" />
		<!-- google analytics -->
		
			<script type="text/javascript">
				var _gaq = _gaq || [];
				_gaq.push(['_setAccount', 'UA-60912543-1']);
				_gaq.push(['_trackPageview']);
				(function() {
				var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
				ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
				var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
				})();
			</script>
		
		<!-- google analytics -->

	</head>
	<body class="ACCUEIL">
		<div id="cookies" style="display: none;">
			<div>
    		<p>En poursuivant votre navigation sur ce site, vous acceptez l'utilisation de cookies pour vous proposer des services et offres adaptés à vos centres d'intérêt.</p>
		      		      <a href="/Pages-annexes/Page-cookie" class="" id="cookies-more">En savoir plus</a>
	          	          <input type="submit" id="cookies-accept" name="cookies-accept" value="OK" title="OK" class="OK"/>
	        </div>
	    </div>
	    
			<script type="text/javascript">
				new CookiesBanner(function(){});
			</script>
		
		<div id="skiptocontent"><a href="#COLONNE-L">Aller au contenu de la page</a></div>
		<div id="CONTAINER">

		<div id="BANDEAU">

	
	
	
<div id="BANDEAU-left">
	<div id="BANDEAU-logo">
		<a title="TCL-SYTRAL, Retourner à l'accueil" href="/">
			<img src="/extension/tcl/design/fr/images/TCL-SYTRAL-logo.png" alt="TCL - SYTRAL" width="107" height="82"  />
		</a>
	</div>
	<div id="BANDEAU-menu-sec">
		<ul class="MENU">
			<li class="espace-emploi"><a href="/Espace-Emploi">emploi</a></li>
			<li class="espace-entreprise"><a href="/Espace-Entreprise">entreprise</a></li>
							<li class="espace-presse"><a href="/Espace-Presse/Les-communiques?theme=all">presse</a></li>
					</ul>
		

	</div>
	<div id="BANDEAU-baseline">
		Partout, pour tous, <strong>il y a TCL</strong>
	</div>
						
<div id="BANDEAU-menu">
	<ul id="dropdown" class="sf-menu">
	                                                				<li><a href="#" >
                                                    Me déplacer
                            					</a>

                                                                
                    
                    						<ul>
                                                            									<li><a href="/Me-deplacer/Itineraires" >Itinéraires</a></li>
                                                                                            									<li><a href="/Me-deplacer/Toutes-les-lignes" >Toutes les lignes</a></li>
                                                                                            									<li><a href="/Me-deplacer/Infos-trafic" >Infos trafic</a></li>
                                                                                            									<li><a href="/Me-deplacer/Plans-du-reseau" >Plans du réseau</a></li>
                                                                                            									<li><a href="/Me-deplacer/Alerte-accessibilite" >Alerte accessibilité</a></li>
                                                            						</ul>
                                        
				</li>
                                                            				<li><a href="#" >
                                                    Tarifs
                            					</a>

                                                                
                    
                    						<ul>
                                                            									<li><a href="/Tarifs/Les-points-de-vente" >Les points de vente</a></li>
                                                                                            									<li><a href="/Tarifs/S-abonner" >S'abonner</a></li>
                                                                                            									<li><a href="/Tarifs/Tickets" >Tickets</a></li>
                                                                                            									<li><a href='https://e-tecely.tcl.fr' target="_blank" >Agence en ligne</a></li>
                                                                                            									<li><a href="/Tarifs/Fraude-et-amendes" >Fraude et amendes</a></li>
                                                            						</ul>
                                        
				</li>
                                                            				<li><a href="#" >
                                                    Découvrir TCL
                            					</a>

                                                                
                    
                    						<ul>
                                                            									<li><a href="/Decouvrir-TCL/Nouveau-client-TCL" >Nouveau client TCL</a></li>
                                                                                            									<li><a href="/Decouvrir-TCL/Le-reseau" >Le réseau</a></li>
                                                                                            									<li><a href="/Decouvrir-TCL/Tous-les-services-TCL" >Tous les services TCL</a></li>
                                                                                            									<li><a href="/Decouvrir-TCL/Accessibilite-du-reseau" >Accessibilité du réseau </a></li>
                                                            						</ul>
                                        
				</li>
                    	</ul>
</div></div>
	 

	
	
	
<div id="BANDEAU-right">

	<div id="BLOC-mon-TCL">
		<div class="BLOC-mon-TCL-content">
			<a id="toogle_bloc_login" href='#' class="mon-tcl">Mon <span>TCL</span></a>
			<div >

			</div>
			

			<div id="divAccountLayerPlaceholder" class="divContent" tabindex="-1">

    <div class="divClose">Fermer X</div>

    <div>
        <!-- User logged in -->
                <h3>Déjà membre ?</h3>
        	<form method="post" action="/user/login" id="connect_page_accueil_form">
		<div id="BLOC-login-over">
			<div class="BLOC-login-over-content">
				<div id="LOGIN-over">
										<div class="COL-over-login">
						<label for="Login" class="LABEL">Adresse email</label>
					    <input class="INPUT-saisie" type="text" name="Login" id="id1" value="" tabindex="1" size="7"/>
					</div> 
					<div class="COL-over-mdp">   
						<label for="MotPasse" class="LABEL">Mot de passe</label>      
					    <input class="INPUT-saisie" type="password" name="Password" id="id2" value="" tabindex="1" size="7"/>
					 	<a href="/user/forgotpassword" class="LIEN">Email ou mot de passe oublié</a>
					</div>
					 <input type="hidden" name="RedirectURI" value="/Mon-TCL" />
					 <input type="submit" name="LoginButton" accesskey="1" title="Valider" class="INPUT-valider" value="Valider"/> 
										<div>
					        <input type="checkbox" tabindex="1" name="Cookie" id="id3" />
					        <label for="id3" style="display:inline;font-size: 0.8em;font-weight: normal;">Se souvenir de moi</label>
					</div>
										<div class="CLEAR-both"></div>
				</div>
			</div>
		</div>

			</form>
	<script>
		var act = $('#connect_page_accueil_form').attr('action');
		var httpsAct = "https://" + window.location.hostname + act;
        $('#connect_page_accueil_form').attr('action', httpsAct);
	</script>                <hr>
                    <h3>Pas encore Membre ?</h3>
            <p>Vous êtes un utilisateur régulier du site TCL.fr ?</p>
            <a class="INPUT-subscribe" href="/user/register" class="LIEN-02"
               title="Créer mon compte">Créer mon compte</a>
        
    </div>
</div>
<script>
    
    $(document).ready(function () {
        $('#toogle_bloc_login').click(function () {
            $('#divAccountLayerPlaceholder').toggle('fade');
        });

        $('.divClose').on('click', function () {
            $('#divAccountLayerPlaceholder').fadeOut();
        });

        // force https on user edit form
        var action = $('#user_edit_form').attr('action');
        var httpsAct = "https://" + window.location.hostname + action;
        $('#user_edit_form').attr('action', httpsAct);
    });
    
</script>			
<div class="FLAG-dropdown">

    <button id="toggle-language"><img src="/extension/tcl/design/fr/images/flag-fre-FR.png"/> FR</button>
    <div class="FLAG">
                                                                                                        <a href="/en">
                                        <img src="/extension/tcl/design/fr/images/flag-eng-GB.png"/> GB
                                    </a>
                                                                                                <a href="/es">
                                        <img src="/extension/tcl/design/fr/images/flag-esl-ES.png"/> ES
                                    </a>
                                                                                                <a href="/it">
                                        <img src="/extension/tcl/design/fr/images/flag-ita-IT.png"/> IT
                                    </a>
                                                                                                <a href="/de">
                                        <img src="/extension/tcl/design/fr/images/flag-ger-DE.png"/> DE
                                    </a>
                                                                </div>
</div><script>
    
    $(document).ready(function () {
        $('#toggle-language').on('click', function(){
            $(".FLAG").toggle();
        });
    });
    
</script>
		</div>

	</div>
	<div class="CLEAR-both"></div>

	<div id="BLOC-recherche">
		            
<form action="/Recherche" method="get" id="Search_form_accueil_toolbar">
	<div class="INPUT-search">
		<div>
			<label for="SearchText-accueil-toolbar">Rechercher sur le site</label>
			<input name="SearchText" type="text" class="INPUT-saisie" value="" size="12" id="SearchText-accueil-toolbar"/>
		</div>
	</div>
	<input type="submit" accesskey="1" value="Trouver" title="Valider la recherche" class="INPUT-valider" />
	</form>
<script type="text/javascript">

	$(document).ready(function() {
		$('#Search').val("Rechercher sur le site");
		
	});
	$('#Search').focus(function() {
		$('#Search').val('');
	});

</script>    
	    	</div>
</div>
	 

	<div class="CLEAR-both"></div>
	<!-- <div id="BANDEAU-sous-menu">-->
		
	<!-- </div>-->
	<script type="text/javascript">
	
	/*************************************Menu déroulant*****************************************/
	$('.sf-menu').addClass("sf-menu-jsactiv");
	$(document).ready(function() {
		// initialise plugin
		var example = $('#dropdown').superfish({
			cssArrows:     false
		});

		// buttons to demonstrate Superfish's public methods
		$('.destroy').on('click', function(){
			example.superfish('destroy');
		});

		$('.init').on('click', function(){
			example.superfish();
		});

		$('.open').on('click', function(){
			example.children('li:first').superfish('show');
		});

		$('.close').on('click', function(){
			example.children('li:first').superfish('hide');
		});	
	});
	/********************************************************************************************/
	
	</script>
</div><!-- FIN BANDEAU -->
																																	
					<div id="CONTENT-pages">
			        	            
    <div id="ARIANE">
    	    		    				    				        				        						            <a href="/">Accueil</a>
					        				        				    				        				        						            <a href="/Mon-TCL">Mon TCL</a>
					        				        				    				        				            <strong>Service Attestation d abonnement</strong>
				        				    		    				    </div>
        				<!-- Bon plan sur la page (col L + col R) -->
		
		
				<div id="COLONNE-L">



					<h1 class="H1-ligne">    Service Attestation d abonnement</h1>	
<a name="eztoc17591584_1" id="eztoc17591584_1"></a><h2>Formulaire de demande d'attestation de paiement</h2>	
	<form method="post" enctype="multipart/form-data" action="/Mon-TCL/Service-Attestation-d-abonnement2">
								<input type="hidden" id="classe_id" name="classe_id" value="132"/>
											
				 
			    	<!-- On créé le titre du formulaire -->
						 
						<input type="hidden" class="INPUT-saisie" id="titre" name="titre" value="Demande_attestationanonymous020520190327"/>
															    		
																<fieldset><legend>Votre demande</legend>
					<table cellspacing="0" cellpadding="0" summary="" class="TABLE-form">
				
				 
				    	<tr>
							<td class="COL-1">
					    										*
												    	<label>Email
					    				    	
					    	</label>
					    </td>
						<td class="COL-2">
											    			<input type="text" class="INPUT-saisie" id="email" name="email" value="pppp@go.com"/>
				    						    	</td>
					</tr>
				    		
											
								    	<tr>
							<td class="COL-1">
						    											*
														    	<label>Date de naissance
						    								    								    	</label>
						    </td>
							<td class="COL-2">
																								 
						    	<label for="jour_date_naissance">Jour</label>					  
						    	<select id="jour_date_naissance" name="jour_date_naissance" class="INPUT-date" >
						            <option value=""></option>
							        							        								        		<option value="01" >1</option>
							            							        							        								        		<option value="02" >2</option>
							            							        							        								        		<option value="03" >3</option>
							            							        							        								        		<option value="04" >4</option>
							            							        							        								        		<option value="05" >5</option>
							            							        							        								        		<option value="06" >6</option>
							            							        							        								        		<option value="07" >7</option>
							            							        							        								        		<option value="08" >8</option>
							            							        							        								        		<option value="09" >9</option>
							            							        							        								            	<option value="10" >10</option>
							        								        							        								            	<option value="11" >11</option>
							        								        							        								            	<option value="12" >12</option>
							        								        							        								            	<option value="13" >13</option>
							        								        							        								            	<option value="14" >14</option>
							        								        							        								            	<option value="15" >15</option>
							        								        							        								            	<option value="16" >16</option>
							        								        							        								            	<option value="17" >17</option>
							        								        							        								            	<option value="18" >18</option>
							        								        							        								            	<option value="19" >19</option>
							        								        							        								            	<option value="20" >20</option>
							        								        							        								            	<option value="21" >21</option>
							        								        							        								            	<option value="22"  selected="selected">22</option>
							        								        							        								            	<option value="23" >23</option>
							        								        							        								            	<option value="24" >24</option>
							        								        							        								            	<option value="25" >25</option>
							        								        							        								            	<option value="26" >26</option>
							        								        							        								            	<option value="27" >27</option>
							        								        							        								            	<option value="28" >28</option>
							        								        							        								            	<option value="29" >29</option>
							        								        							        								            	<option value="30" >30</option>
							        								        							        								            	<option value="31" >31</option>
							        								        						        </select>
						         
						    	<label for="mois_date_naissance">Mois</label>
						    	<select id="mois_date_naissance" name="mois_date_naissance" class="INPUT-date">
						            <option value=""></option>
							        							        								            	<option value="01"  >1</option>
							        								        							        								            	<option value="02"  >2</option>
							        								        							        								            	<option value="03"  selected="selected" >3</option>
							        								        							        								            	<option value="04"  >4</option>
							        								        							        								            	<option value="05"  >5</option>
							        								        							        								            	<option value="06"  >6</option>
							        								        							        								            	<option value="07"  >7</option>
							        								        							        								            	<option value="08"  >8</option>
							        								        							        								            	<option value="09"  >9</option>
							        								        							        								        		<option value="10"   >10</option>
							        								        							        								        		<option value="11"   >11</option>
							        								        							        								        		<option value="12"   >12</option>
							        								        						        </select>
						    	 
						    	<label for="annee_date_naissance">Année</label>
								<select id="annee_date_naissance" name="annee_date_naissance" class="INPUT-date">
						            <option value=""></option>
							        							             <option value="2019" >2019</option>
							        							             <option value="2018" >2018</option>
							        							             <option value="2017" >2017</option>
							        							             <option value="2016" >2016</option>
							        							             <option value="2015" >2015</option>
							        							             <option value="2014" >2014</option>
							        							             <option value="2013" >2013</option>
							        							             <option value="2012" >2012</option>
							        							             <option value="2011" >2011</option>
							        							             <option value="2010" >2010</option>
							        							             <option value="2009" >2009</option>
							        							             <option value="2008" >2008</option>
							        							             <option value="2007" >2007</option>
							        							             <option value="2006" >2006</option>
							        							             <option value="2005" >2005</option>
							        							             <option value="2004" >2004</option>
							        							             <option value="2003" >2003</option>
							        							             <option value="2002" >2002</option>
							        							             <option value="2001" >2001</option>
							        							             <option value="2000" >2000</option>
							        							             <option value="1999" >1999</option>
							        							             <option value="1998" >1998</option>
							        							             <option value="1997" >1997</option>
							        							             <option value="1996" >1996</option>
							        							             <option value="1995" >1995</option>
							        							             <option value="1994" >1994</option>
							        							             <option value="1993" >1993</option>
							        							             <option value="1992" >1992</option>
							        							             <option value="1991" >1991</option>
							        							             <option value="1990" >1990</option>
							        							             <option value="1989" >1989</option>
							        							             <option value="1988" >1988</option>
							        							             <option value="1987" >1987</option>
							        							             <option value="1986" >1986</option>
							        							             <option value="1985"  selected="selected">1985</option>
							        							             <option value="1984" >1984</option>
							        							             <option value="1983" >1983</option>
							        							             <option value="1982" >1982</option>
							        							             <option value="1981" >1981</option>
							        							             <option value="1980" >1980</option>
							        							             <option value="1979" >1979</option>
							        							             <option value="1978" >1978</option>
							        							             <option value="1977" >1977</option>
							        							             <option value="1976" >1976</option>
							        							             <option value="1975" >1975</option>
							        							             <option value="1974" >1974</option>
							        							             <option value="1973" >1973</option>
							        							             <option value="1972" >1972</option>
							        							             <option value="1971" >1971</option>
							        							             <option value="1970" >1970</option>
							        							             <option value="1969" >1969</option>
							        							             <option value="1968" >1968</option>
							        							             <option value="1967" >1967</option>
							        							             <option value="1966" >1966</option>
							        							             <option value="1965" >1965</option>
							        							             <option value="1964" >1964</option>
							        							             <option value="1963" >1963</option>
							        							             <option value="1962" >1962</option>
							        							             <option value="1961" >1961</option>
							        							             <option value="1960" >1960</option>
							        							             <option value="1959" >1959</option>
							        							             <option value="1958" >1958</option>
							        							             <option value="1957" >1957</option>
							        							             <option value="1956" >1956</option>
							        							             <option value="1955" >1955</option>
							        							             <option value="1954" >1954</option>
							        							             <option value="1953" >1953</option>
							        							             <option value="1952" >1952</option>
							        							             <option value="1951" >1951</option>
							        							             <option value="1950" >1950</option>
							        							             <option value="1949" >1949</option>
							        							             <option value="1948" >1948</option>
							        							             <option value="1947" >1947</option>
							        							             <option value="1946" >1946</option>
							        							             <option value="1945" >1945</option>
							        							             <option value="1944" >1944</option>
							        							             <option value="1943" >1943</option>
							        							             <option value="1942" >1942</option>
							        							             <option value="1941" >1941</option>
							        							             <option value="1940" >1940</option>
							        							             <option value="1939" >1939</option>
							        							             <option value="1938" >1938</option>
							        							             <option value="1937" >1937</option>
							        							             <option value="1936" >1936</option>
							        							             <option value="1935" >1935</option>
							        							             <option value="1934" >1934</option>
							        							             <option value="1933" >1933</option>
							        							             <option value="1932" >1932</option>
							        							             <option value="1931" >1931</option>
							        							             <option value="1930" >1930</option>
							        							             <option value="1929" >1929</option>
							        							             <option value="1928" >1928</option>
							        							             <option value="1927" >1927</option>
							        							             <option value="1926" >1926</option>
							        							             <option value="1925" >1925</option>
							        							             <option value="1924" >1924</option>
							        							             <option value="1923" >1923</option>
							        							             <option value="1922" >1922</option>
							        							             <option value="1921" >1921</option>
							        							             <option value="1920" >1920</option>
							        							             <option value="1919" >1919</option>
							        						        </select>
														</td>
						</tr>
				    		
											
										<tr>
							<td class="COL-1">
																<label for="type_carte">Type de carte
									</label>
							</td>
															<td class="COL-2">
									<select id="type_carte" name="type_carte" class="INPUT-type-carte">
																					<option value="0" selected="selected">
												Carte Técély
											</option>
																			</select>
								</td>
													</tr>
				    		
											
				 
			    					    	<tr>
						    <td class="COL-1">
						    										*
														    	<label for="num_tecely">Numero de carte
			    				<span class="BLOC-alerte">Un numéro de carte correct est obligatoire</span></label>
						    </td>
				    									    <td class="COL-2">
							    								    		<input type="text"  class="INPUT-saisie ALERTE"  id="num_tecely" name="num_tecely" value="plop"/>
							    								    </td>
													 </tr>
									    		
											
								    	<tr>
							<td class="COL-1">
						    											*
														    	<label>Date début
						    								    								    	</label>
						    </td>
							<td class="COL-2">
															 
														    		<input type="hidden" class="INPUT-saisie" id="jour_date_debut" name="jour_date_debut" value="01"/>
						        						         
						    	<label for="mois_date_debut">Mois</label>
						    	<select id="mois_date_debut" name="mois_date_debut" class="INPUT-date-mois">
						            <option value=""></option>
							        							        								            	<option value="01"  >Janvier</option>
							        								        							        								            	<option value="02"  >Février</option>
							        								        							        								            	<option value="03"  >Mars</option>
							        								        							        								            	<option value="04"  >Avril</option>
							        								        							        								            	<option value="05"  selected="selected" >Mai</option>
							        								        							        								            	<option value="06"  >Juin</option>
							        								        							        								            	<option value="07"  >Juillet</option>
							        								        							        								            	<option value="08"  >Août</option>
							        								        							        								            	<option value="09"  >Septembre</option>
							        								        							        								        		<option value="10"   >Octobre</option>
							        								        							        								        		<option value="11"   >Novembre</option>
							        								        							        								        		<option value="12"   >Décembre</option>
							        								        						        </select>
						    	 
						    	<label for="annee_date_debut">Année</label>
								<select id="annee_date_debut" name="annee_date_debut" class="INPUT-date">
						            <option value=""></option>
							        							             <option value="2019"  selected="selected">2019</option>
							        							             <option value="2018" >2018</option>
							        						        </select>
						    							</td>
						</tr>
				    		
											
								    	<tr>
							<td class="COL-1">
						    											*
														    	<label>Date de fin
						    								    								    	</label>
						    </td>
							<td class="COL-2">
															 
														        	<input type="hidden" class="INPUT-saisie" id="jour_date_fin" name="jour_date_fin" value="28"/>
						        						         
						    	<label for="mois_date_fin">Mois</label>
						    	<select id="mois_date_fin" name="mois_date_fin" class="INPUT-date-mois">
						            <option value=""></option>
							        							        								            	<option value="01"  >Janvier</option>
							        								        							        								            	<option value="02"  >Février</option>
							        								        							        								            	<option value="03"  >Mars</option>
							        								        							        								            	<option value="04"  >Avril</option>
							        								        							        								            	<option value="05"  selected="selected" >Mai</option>
							        								        							        								            	<option value="06"  >Juin</option>
							        								        							        								            	<option value="07"  >Juillet</option>
							        								        							        								            	<option value="08"  >Août</option>
							        								        							        								            	<option value="09"  >Septembre</option>
							        								        							        								        		<option value="10"   >Octobre</option>
							        								        							        								        		<option value="11"   >Novembre</option>
							        								        							        								        		<option value="12"   >Décembre</option>
							        								        						        </select>
						    	 
						    	<label for="annee_date_fin">Année</label>
								<select id="annee_date_fin" name="annee_date_fin" class="INPUT-date">
						            <option value=""></option>
							        							             <option value="2019"  selected="selected">2019</option>
							        							             <option value="2018" >2018</option>
							        						        </select>
						    							</td>
						</tr>
				    		
											
						
						</table>
		</fieldset>
			<div class="ALIGN-center MARGIN-top-15px">
				<div class="bottom-mon-tcl">
					<input type="submit" id="envoyer" name="envoyer" value="Envoyer" title="Envoyer" />
				</div>
			</div>		
			</form>

<p><b>Tous les champs sont obligatoires !</b></p><p>&nbsp;</p><a name="eztoc17591585_0_1" id="eztoc17591585_0_1"></a><h3><b>Bien remplir sa demande d’attestation </b></h3><p><b>&nbsp;</b></p><p><b>Numéro de carte TECELY</b> : Ce numéro figure au dos de votre carte TECELY et comporte 12 numéros (e<i>xemple</i>&nbsp;: <b>021231456987)</b></p><p><b>Date de début et date de fin&nbsp;</b>: Ces 2 dates correspondent à la période pour laquelle vous souhaitez obtenir une attestation de paiement.</p><p>&nbsp;</p><p>Exemple&nbsp;:</p><p>&nbsp;</p><p>Vous souhaitez une attestation de paiement pour <b>le mois de juillet 2017</b>.</p><p>Il faudra sélectionner comme <b>date de début</b> «&nbsp;Juillet 2017 » &nbsp;et comme <b>date de fin</b> «&nbsp;Juillet 2017 ».</p><p>Vous souhaitez une attestation de paiement pour <b>la période du 1er janvier 2017 au 31 juillet 2017</b>.</p>
<ul>

<li>Il faudra sélectionner comme <b>date de début</b> «&nbsp;Janvier 2017 » et comme<b>&nbsp;date de fin</b> «&nbsp;Juillet 2017 ».</li>

</ul>
<p>&nbsp;</p><a name="eztoc17591585_0_2" id="eztoc17591585_0_2"></a><h3><b>A partir de quand je peux demander une attestation&nbsp;?</b></h3><p>&nbsp;</p><p>Pour les abonnements par prélèvement automatique les attestations de paiement sont disponibles à partir du 2 du mois en cours.</p><p>Exemple&nbsp;: Pour obtenir une attestation pour le mois de juillet 2017, il faudra attendre le 2 juilllet 2017 pour faire votre demande d’attestation de paiement.</p><p>&nbsp;</p><a name="eztoc17591585_0_3" id="eztoc17591585_0_3"></a><h3><b>Délai d’obtention de mon attestation de paiement</b></h3><p>&nbsp;</p><p>Vous obtiendrez par mail votre attestation dans un délai de 72h maximum, vérifiez donc bien que votre adresse mail est correctement renseignée.</p>

					<!-- Module Navitia / lignes_a_proximite bloc bas -->
										<!-- Module Navitia / lignes_a_proximite bloc milieu -->
										<!-- Module Google / google_maps -->
										<!-- Formulaire Mes alertes -->
										<!-- Alertes accessibilité -->
					
					 <!-- Magic node : Page mécanisme colonnes - NE PAS TOUCHER -->
					
									</div>

				<div id="COLONNE-R">

					 

					<!-- Bloc Info trafic accueil -->
					
					
					
					<!-- Bloc Services -->
																							<div class="BLOC-col-D">
                                <!-- Bloc Services -->
                                
                                <div id="encart-itineraire" class="nounderline white">
                                    <a href="/Me-deplacer/Itineraires/Mon-trajet">Calculer mon itinéraire</a>
                                </div>


                                    <div>
                                        <div class="contenu mon-trajet-services">
                                            <ul>
                                                                                                                                                                                                        
                                                                                                                                                                                                            
                                                                                                            <li>


<a href="/Me-deplacer/Toutes-les-lignes"
   target="_self"
   title="Accédez au service Horaires Horaires">
		<span class="picto">
			<img src="/var/tcl/storage/images/media/librairie/encart-service/horaires/360-18-fre-FR/Horaires.png"
                 alt="    Horaires"/>
		</span>
        <span class="text-service">
                Horaires        </span>
</a>


</li>
                                                                                                                                                                                                            
                                                                                                            <li>
            <a href="/Me-deplacer/Plans-du-reseau"
       class="append"
       target="_self"
       title="Accédez au service PDF Plans du réseau">
        <span class="type-file">    PDF</span>
    </a>
    
<a href="http://plan-interactif.tcl.fr/"
   target="_blank"
   title="Accédez au service Plan interactif (nouvelle fenêtre)">
		<span class="picto">
			<img src="/var/tcl/storage/images/media/librairie/encart-service/plan-interactif/374-19-fre-FR/Plan-interactif.png"
                 alt="    Plan interactif"/>
		</span>
        <span class="text-service">
                Plan interactif        </span>
</a>


</li>
                                                                                                                                                                                                            
                                                                                                                                                                                                            
                                                                                                            <li>

    
<a href="https://e-tecely.tcl.fr/accueil.do"
   target="_blank"
   title="Accédez au service Agence en ligne Agence en ligne">
		<span class="picto">
			<img src="/var/tcl/storage/images/media/librairie/encart-service/agence-en-ligne/410-17-fre-FR/Agence-en-ligne.png"
                 alt="    Agence en ligne"/>
		</span>
        <span class="text-service">
                Agence en ligne        </span>
</a>


</li>
                                                                                                                                                                                                            
                                                                                                            <li>


<a href="/Tarifs/Fraude-et-amendes"
   target="_self"
   title="Accédez au service Paiement des amendes Paiement des amendes">
		<span class="picto">
			<img src="/var/tcl/storage/images/media/librairie/encart-service/paiement-des-amendes/381-28-fre-FR/Paiement-des-amendes.png"
                 alt="    Paiement des amendes"/>
		</span>
        <span class="text-service">
                Paiement des amendes        </span>
</a>


</li>
                                                                                                                                                                                                            
                                                                                                                                                                                                </ul>
                                            <div class="CLEAR-both"></div>

                                        </div>
                                        <div class="alertes-accessibilite nounderline">
                                                                        <h3><a href="/Me-deplacer/Alerte-accessibilite">Alertes accessibilité <span>(10)</span></a></h3>
                                                                    </div>
                                    </div>
							</div>
																					<!-- Bloc Actu en bref -->
										<!-- Bloc Guide tarifaire -->
					
					 

					<!-- Bloc Info trafic -->
										<!-- Bloc Horaire Arret -->
										<!-- Bloc Alerte Trafic -->
					

					
					
					<!-- Bloc Parc relais de la ligne -->
										<!-- Bloc Actualité de la ligne -->
										<!-- Bloc Carte tecely  -->
					
					
					 
					
					
					
										<!-- Bloc Mes coordonnées -->
										<!-- Bloc Mes alertes -->
										<!-- Bloc Mes newsletters -->
					
										
					 

					<!-- Bloc Alertes accessibilité -->
										<!-- Bloc Fraude -->
									</div>

				<div class="CLEAR-both"></div>
				

				<!-- FIN CONTENT -->
			</div>
		

			
			
			
			

<div id="FOOTER">
	<div id="BLOC-footer">
				 <div class="BLOC">
		 			 				 <h2>    Contacter TCL</h2>
				 
<p class="NUM-tel">Allô TCL&nbsp;<b>04 26 10 12 12</b></p><p class="NUM-tel-info">Prix d’un appel local</p><p>Du lundi au dimanche : 6h / 22h </p>			 		 			 				 <h2><a class="nous-ecrire" href="/Pied-de-page/Nous-ecrire">    Nous écrire</a></h2>
			 		 		 </div>
				
					<div class="BLOC">
										<h2><a href="/Me-deplacer/Itineraires">Me déplacer</a></h2>

													<ul>
																	<li><a href="/Me-deplacer/Itineraires">Itinéraires</a></li>
																							<li><a href="/Me-deplacer/Toutes-les-lignes">Toutes les lignes</a></li>
																							<li><a href="/Me-deplacer/Infos-trafic">Infos trafic</a></li>
																							<li><a href="/Me-deplacer/Plans-du-reseau">Plans du réseau</a></li>
																							<li><a href="/Me-deplacer/Alerte-accessibilite">Alerte accessibilité</a></li>
																</ul>
														</div>
					<div class="BLOC">
										<h2><a href="/Tarifs/Les-points-de-vente">Tarifs</a></h2>

													<ul>
																	<li><a href="/Tarifs/Les-points-de-vente">Les points de vente</a></li>
																							<li><a href="/Tarifs/S-abonner">S'abonner</a></li>
																							<li><a href="/Tarifs/Tickets">Tickets</a></li>
																							<li><a href='https://e-tecely.tcl.fr' target="_blank" >Agence en ligne</a></li>
																							<li><a href="/Tarifs/Fraude-et-amendes">Fraude et amendes</a></li>
																</ul>
														</div>
					<div class="BLOC">
										<h2><a href="/Decouvrir-TCL/Nouveau-client-TCL">Découvrir TCL</a></h2>

													<ul>
																	<li><a href="/Decouvrir-TCL/Nouveau-client-TCL">Nouveau client TCL</a></li>
																							<li><a href="/Decouvrir-TCL/Le-reseau">Le réseau</a></li>
																							<li><a href="/Decouvrir-TCL/Tous-les-services-TCL">Tous les services TCL</a></li>
																							<li><a href="/Decouvrir-TCL/Accessibilite-du-reseau">Accessibilité du réseau </a></li>
																</ul>
														</div>
				<div class="BLOC">
										<h2>Suivez nous</h2>

				<ul class="list-inline">
									<li>  <a href=https://twitter.com/tcl_sytral> <img src="/var/tcl/storage/images/pied-de-page/suivez-nous/twitter/32041250-1-fre-FR/Twitter.png" alt='Twitter'>  </a></li>
									<li>  <a href=https://www.instagram.com/instatram_tcl/> <img src="/var/tcl/storage/images/pied-de-page/suivez-nous/instagram/32041299-1-fre-FR/Instagram.png" alt='Instagram'>  </a></li>
									<li>  <a href=https://www.facebook.com/TCL.SYTRAL> <img src="/var/tcl/storage/images/pied-de-page/suivez-nous/facebook/32041326-1-fre-FR/Facebook.png" alt='Facebook'>  </a></li>
								</ul>
						<br>
									<h2>Télécharger l'appli</h2>

			<ul class="list-inline">
									<li>  <a href=https://itunes.apple.com/fr/app/tcl/id579379606?mt=8> <img src="/var/tcl/storage/images/pied-de-page/telecharger-l-appli/iphone/32041182-1-fre-FR/Iphone.png" alt='Iphone'>  </a></li>
									<li>  <a href=https://play.google.com/store/apps/details?id=com.micropole.android.tcl_mobile> <img src="/var/tcl/storage/images/pied-de-page/telecharger-l-appli/android/32041239-1-fre-FR/Android.png" alt='Android'>  </a></li>
											</ul>
		</div>

		<a href="#" class="LIEN-haut">HAUT</a>
		<div class="CLEAR-both"></div>
	</div>
	<div id="FOOTER-END">
		
<div  class="logo-footer-end" >
	<p>						


    
        
                                
                        <img              src="/var/tcl/storage/images/media/images/logo-tcl-footer/32042452-1-fre-FR/logo-tcl-footer.png"             width="50"
             height="38"              alt=""
             />
            


    
    			   </p>
</div><div  class="slogan-footer" >
	<p><i>Bougez, Vivez, Aimez !</i></p>
</div><div  class="liens-footer-end" >
	
<ul>

<li>TCL.fr ©2019</li>

<li><a href="http://www.tcl.fr/content/download/1249414/17216300/version/3/file/CHARTE+CLIENTS+TCL.PDF" target="_self">Charte client</a></li>

<li><a href="/Pied-de-page/Infos/Mentions-legales" title="Mentions légales" target="_self">Mentions légales</a></li>

<li><a href="/Pied-de-page/Infos/Plan-du-site" title="Plan du site" target="_self">Plan du Site</a></li>

<li><a href="http://tcl.fr/Pied-de-page/Infos/Participer-au-panel-TCL" target="_self">Le panel TCL</a></li>

</ul>

</div><div  class="certification-footer-end" >
	<p><a href="http://www.afnor.org" title="Accéder au site AFNOR-certification (nouvelle fenêtre)" target="_blank" class="lien-afnor">						


    
        
                                
                        <img              src="/var/tcl/storage/images/media/images/afnor-certification/4155125-1-fre-FR/AFNOR-certification.png"             width="114"
             height="35"              alt="AFNOR-certification"
             />
            


    
    			   </a><a href="http://www.accessiweb.org/index.php/rapport_de_labellisation/items/tcl_2012.html" title="Label AccessiWeb Argent 2012 - Transports en Commun Lyonnais (rapport de labellisation) (nouvelle fenêtre)" target="_blank" class="lien-accessibilite">						


    
        
                                
                        <img              src="/var/tcl/storage/images/media/images/label-accessiweb-argent-2012-transports-en-commun-lyonnais-rapport-de-labellisation/4155128-1-fre-FR/Label-AccessiWeb-Argent-2012-Transports-en-Commun-Lyonnais-rapport-de-labellisation.png"             width="114"
             height="35"              alt="Label AccessiWeb Argent 2012 - Transports en Commun Lyonnais (rapport de labellisation)"
             />
            


    
    			   </a></p>
</div>	</div>
</div><!-- fin FOOTER -->
		</div> <!-- fin CONTAINER -->

																	<!-- eStat -->
		
			<script type="text/javascript">
				<!--
				var _PJS=0;
				//-->
			</script>
			<script type="text/javascript" src="//prof.estat.com/js/276076187740.js"></script>
			<script type="text/javascript">
				<!--
				if(_PJS){
					eStat_id.serial("276076187740");
					eStat_id.master("218018158547");
					eStat_id.gp_pg_mq("index.asp");
					eStat_id.pg_mq("index.asp");
					eStat_tag.post("m");
				}
				//-->
			</script>
			<noscript>
				<img src="http://prof.estat.com/m/web/276076187740?g=218018158547&amp;c=index.asp&amp;p=index.asp&amp;st=0&amp;sjs=0" style="width:1; height:1;" alt="eStat Pro"/>
			</noscript>
			<!-- /eStat -->
		
	</body>
</html>
`

const okHtmlOutput = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
	"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="fr-FR" lang="fr-FR">
					<head>
		                            				
                	<title>TCL - Service Attestation d abonnement</title>
    
    
    
    				<meta name="keywords" content="Service Attestation d abonnement" />
			    	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
            <meta name="Content-Type" content="text/html; charset=utf-8" />

            <meta name="Content-language" content="fr-FR" />

                    <meta name="author" content="Micropole" />
    
                <meta name="copyright" content="Tcl.fr" />
    
                <meta name="description" content="TCL, réseau des Transports en Commun Lyonnais. Plan des lignes de métro, bus, tramway et funiculaire sur Lyon, Villeurbanne et l'agglomération lyonnaise ; horaires des bus, trams et métros ; trajets sur-mesure : calculer un itinéraire ; tarif des tickets de bus, carnets, abonnements, cartes Técély" />
    
                <meta name="keywords" content="TCL, transports en commun lyonnais, bus rhône, tram lyon, métro ville lyon, funiculaire lyon, tramway lyon, plan lignes, lignes métro, lignes bus, lignes tramways, lignes funiculaire, transports en commun lyon, transports en commun Villeurbanne, transports en commun Rhône, horaires bus, horaires trams, horaires métros, calcul trajets, calcul itinéraire, tarif tickets bus, tarif abonnement, tarif carte abonnement, tarif carte Técély, bons plans lyon" />
    
        <meta name="MSSmartTagsPreventParsing" content="TRUE" />
		<link rel="stylesheet" type="text/css" href="/var/tcl/cache/public/stylesheets/6f4a838f23e9676c03c193a31c890c93_1554111355_screen.css" media="screen" />


<link rel="stylesheet" type="text/css" href="/var/tcl/cache/public/stylesheets/96a690e8a6d5a55c1f81991515a26111_1554111355_print.css" media="print" />



        <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,400,700,800" rel="stylesheet">

<!--[if IE]>
		<link href="/extension/tcl/design/fr/stylesheets/ie-tous.css" rel="stylesheet" type="text/css" media="screen" />
		<link href="/extension/tcl/design/fr/stylesheets/ie-tous_print.css" rel="stylesheet" type="text/css" media="print" />
<![endif]-->

<!--[if lt IE 7]>
		<link href="/extension/tcl/design/fr/stylesheets/ie6.css" rel="stylesheet" type="text/css" media="screen" />
<![endif]-->

<!--[if IE 7]>
		<link href="/extension/tcl/design/fr/stylesheets/ie7.css" rel="stylesheet" type="text/css" media="screen" />
<![endif]-->
<!--[if IE 8]>
		<link href="/extension/tcl/design/fr/stylesheets/ie8.css" rel="stylesheet" type="text/css" media="screen" />
<![endif]-->
<!--[if gt IE 7]>
		<link href="/extension/tcl/design/fr/stylesheets/ie-tous.css" rel="stylesheet" type="text/css" media="screen" />
<![endif]-->
		<script type="text/javascript" src="/var/tcl/cache/public/javascript/06808eb35470cc7c2a9fc0609802e2e6_1554111354.js" charset="utf-8"></script>


<script type="text/javascript" src="/var/tcl/cache/public/javascript/37cff028e4be8809d6322fe585f6f543_1554111355.js" charset="utf-8"></script>


		
				
		<link rel="Shortcut icon" href="/extension/tcl/design/fr/images/favicon.ico" type="image/x-icon" />
		<!-- google analytics -->
		
			<script type="text/javascript">
				var _gaq = _gaq || [];
				_gaq.push(['_setAccount', 'UA-60912543-1']);
				_gaq.push(['_trackPageview']);
				(function() {
				var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
				ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
				var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
				})();
			</script>
		
		<!-- google analytics -->

	</head>
	<body class="ACCUEIL">
		<div id="cookies" style="display: none;">
			<div>
    		<p>En poursuivant votre navigation sur ce site, vous acceptez l'utilisation de cookies pour vous proposer des services et offres adaptés à vos centres d'intérêt.</p>
		      		      <a href="/Pages-annexes/Page-cookie" class="" id="cookies-more">En savoir plus</a>
	          	          <input type="submit" id="cookies-accept" name="cookies-accept" value="OK" title="OK" class="OK"/>
	        </div>
	    </div>
	    
			<script type="text/javascript">
				new CookiesBanner(function(){});
			</script>
		
		<div id="skiptocontent"><a href="#COLONNE-L">Aller au contenu de la page</a></div>
		<div id="CONTAINER">

		<div id="BANDEAU">

	
	
	
<div id="BANDEAU-left">
	<div id="BANDEAU-logo">
		<a title="TCL-SYTRAL, Retourner à l'accueil" href="/">
			<img src="/extension/tcl/design/fr/images/TCL-SYTRAL-logo.png" alt="TCL - SYTRAL" width="107" height="82"  />
		</a>
	</div>
	<div id="BANDEAU-menu-sec">
		<ul class="MENU">
			<li class="espace-emploi"><a href="/Espace-Emploi">emploi</a></li>
			<li class="espace-entreprise"><a href="/Espace-Entreprise">entreprise</a></li>
							<li class="espace-presse"><a href="/Espace-Presse/Les-communiques?theme=all">presse</a></li>
					</ul>
		

	</div>
	<div id="BANDEAU-baseline">
		Partout, pour tous, <strong>il y a TCL</strong>
	</div>
						
<div id="BANDEAU-menu">
	<ul id="dropdown" class="sf-menu">
	                                                				<li><a href="#" >
                                                    Me déplacer
                            					</a>

                                                                
                    
                    						<ul>
                                                            									<li><a href="/Me-deplacer/Itineraires" >Itinéraires</a></li>
                                                                                            									<li><a href="/Me-deplacer/Toutes-les-lignes" >Toutes les lignes</a></li>
                                                                                            									<li><a href="/Me-deplacer/Infos-trafic" >Infos trafic</a></li>
                                                                                            									<li><a href="/Me-deplacer/Plans-du-reseau" >Plans du réseau</a></li>
                                                                                            									<li><a href="/Me-deplacer/Alerte-accessibilite" >Alerte accessibilité</a></li>
                                                            						</ul>
                                        
				</li>
                                                            				<li><a href="#" >
                                                    Tarifs
                            					</a>

                                                                
                    
                    						<ul>
                                                            									<li><a href="/Tarifs/Les-points-de-vente" >Les points de vente</a></li>
                                                                                            									<li><a href="/Tarifs/S-abonner" >S'abonner</a></li>
                                                                                            									<li><a href="/Tarifs/Tickets" >Tickets</a></li>
                                                                                            									<li><a href='https://e-tecely.tcl.fr' target="_blank" >Agence en ligne</a></li>
                                                                                            									<li><a href="/Tarifs/Fraude-et-amendes" >Fraude et amendes</a></li>
                                                            						</ul>
                                        
				</li>
                                                            				<li><a href="#" >
                                                    Découvrir TCL
                            					</a>

                                                                
                    
                    						<ul>
                                                            									<li><a href="/Decouvrir-TCL/Nouveau-client-TCL" >Nouveau client TCL</a></li>
                                                                                            									<li><a href="/Decouvrir-TCL/Le-reseau" >Le réseau</a></li>
                                                                                            									<li><a href="/Decouvrir-TCL/Tous-les-services-TCL" >Tous les services TCL</a></li>
                                                                                            									<li><a href="/Decouvrir-TCL/Accessibilite-du-reseau" >Accessibilité du réseau </a></li>
                                                            						</ul>
                                        
				</li>
                    	</ul>
</div></div>
	 

	
	
	
<div id="BANDEAU-right">

	<div id="BLOC-mon-TCL">
		<div class="BLOC-mon-TCL-content">
			<a id="toogle_bloc_login" href='#' class="mon-tcl">Mon <span>TCL</span></a>
			<div >

			</div>
			

			<div id="divAccountLayerPlaceholder" class="divContent" tabindex="-1">

    <div class="divClose">Fermer X</div>

    <div>
        <!-- User logged in -->
                <h3>Déjà membre ?</h3>
        	<form method="post" action="/user/login" id="connect_page_accueil_form">
		<div id="BLOC-login-over">
			<div class="BLOC-login-over-content">
				<div id="LOGIN-over">
										<div class="COL-over-login">
						<label for="Login" class="LABEL">Adresse email</label>
					    <input class="INPUT-saisie" type="text" name="Login" id="id1" value="" tabindex="1" size="7"/>
					</div> 
					<div class="COL-over-mdp">   
						<label for="MotPasse" class="LABEL">Mot de passe</label>      
					    <input class="INPUT-saisie" type="password" name="Password" id="id2" value="" tabindex="1" size="7"/>
					 	<a href="/user/forgotpassword" class="LIEN">Email ou mot de passe oublié</a>
					</div>
					 <input type="hidden" name="RedirectURI" value="/Mon-TCL" />
					 <input type="submit" name="LoginButton" accesskey="1" title="Valider" class="INPUT-valider" value="Valider"/> 
										<div>
					        <input type="checkbox" tabindex="1" name="Cookie" id="id3" />
					        <label for="id3" style="display:inline;font-size: 0.8em;font-weight: normal;">Se souvenir de moi</label>
					</div>
										<div class="CLEAR-both"></div>
				</div>
			</div>
		</div>

			</form>
	<script>
		var act = $('#connect_page_accueil_form').attr('action');
		var httpsAct = "https://" + window.location.hostname + act;
        $('#connect_page_accueil_form').attr('action', httpsAct);
	</script>                <hr>
                    <h3>Pas encore Membre ?</h3>
            <p>Vous êtes un utilisateur régulier du site TCL.fr ?</p>
            <a class="INPUT-subscribe" href="/user/register" class="LIEN-02"
               title="Créer mon compte">Créer mon compte</a>
        
    </div>
</div>
<script>
    
    $(document).ready(function () {
        $('#toogle_bloc_login').click(function () {
            $('#divAccountLayerPlaceholder').toggle('fade');
        });

        $('.divClose').on('click', function () {
            $('#divAccountLayerPlaceholder').fadeOut();
        });

        // force https on user edit form
        var action = $('#user_edit_form').attr('action');
        var httpsAct = "https://" + window.location.hostname + action;
        $('#user_edit_form').attr('action', httpsAct);
    });
    
</script>			
<div class="FLAG-dropdown">

    <button id="toggle-language"><img src="/extension/tcl/design/fr/images/flag-fre-FR.png"/> FR</button>
    <div class="FLAG">
                                                                                                        <a href="/en">
                                        <img src="/extension/tcl/design/fr/images/flag-eng-GB.png"/> GB
                                    </a>
                                                                                                <a href="/es">
                                        <img src="/extension/tcl/design/fr/images/flag-esl-ES.png"/> ES
                                    </a>
                                                                                                <a href="/it">
                                        <img src="/extension/tcl/design/fr/images/flag-ita-IT.png"/> IT
                                    </a>
                                                                                                <a href="/de">
                                        <img src="/extension/tcl/design/fr/images/flag-ger-DE.png"/> DE
                                    </a>
                                                                </div>
</div><script>
    
    $(document).ready(function () {
        $('#toggle-language').on('click', function(){
            $(".FLAG").toggle();
        });
    });
    
</script>
		</div>

	</div>
	<div class="CLEAR-both"></div>

	<div id="BLOC-recherche">
		            
<form action="/Recherche" method="get" id="Search_form_accueil_toolbar">
	<div class="INPUT-search">
		<div>
			<label for="SearchText-accueil-toolbar">Rechercher sur le site</label>
			<input name="SearchText" type="text" class="INPUT-saisie" value="" size="12" id="SearchText-accueil-toolbar"/>
		</div>
	</div>
	<input type="submit" accesskey="1" value="Trouver" title="Valider la recherche" class="INPUT-valider" />
	</form>
<script type="text/javascript">

	$(document).ready(function() {
		$('#Search').val("Rechercher sur le site");
		
	});
	$('#Search').focus(function() {
		$('#Search').val('');
	});

</script>    
	    	</div>
</div>
	 

	<div class="CLEAR-both"></div>
	<!-- <div id="BANDEAU-sous-menu">-->
		
	<!-- </div>-->
	<script type="text/javascript">
	
	/*************************************Menu déroulant*****************************************/
	$('.sf-menu').addClass("sf-menu-jsactiv");
	$(document).ready(function() {
		// initialise plugin
		var example = $('#dropdown').superfish({
			cssArrows:     false
		});

		// buttons to demonstrate Superfish's public methods
		$('.destroy').on('click', function(){
			example.superfish('destroy');
		});

		$('.init').on('click', function(){
			example.superfish();
		});

		$('.open').on('click', function(){
			example.children('li:first').superfish('show');
		});

		$('.close').on('click', function(){
			example.children('li:first').superfish('hide');
		});	
	});
	/********************************************************************************************/
	
	</script>
</div><!-- FIN BANDEAU -->
																																	
					<div id="CONTENT-pages">
			        	            
    <div id="ARIANE">
    	    		    				    				        				        						            <a href="/">Accueil</a>
					        				        				    				        				        						            <a href="/Mon-TCL">Mon TCL</a>
					        				        				    				        				            <strong>Service Attestation d abonnement</strong>
				        				    		    				    </div>
        				<!-- Bon plan sur la page (col L + col R) -->
		
		
				<div id="COLONNE-L">



					<h1 class="H1-ligne">    Service Attestation d abonnement</h1>	
<a name="eztoc17591584_1" id="eztoc17591584_1"></a><h2>Formulaire de demande d'attestation de paiement</h2>	
	<form method="post" enctype="multipart/form-data" action="/Mon-TCL/Service-Attestation-d-abonnement2">
					<div class="BLOC-simple MARGIN-top-10px MARGIN-bottom-10px">
				<p class="ALIGN-center">Votre demande d'attestation a bien été prise en compte</p>
			</div>
			</form>

<p><b>Tous les champs sont obligatoires !</b></p><p>&nbsp;</p><a name="eztoc17591585_0_1" id="eztoc17591585_0_1"></a><h3><b>Bien remplir sa demande d’attestation </b></h3><p><b>&nbsp;</b></p><p><b>Numéro de carte TECELY</b> : Ce numéro figure au dos de votre carte TECELY et comporte 12 numéros (e<i>xemple</i>&nbsp;: <b>021231456987)</b></p><p><b>Date de début et date de fin&nbsp;</b>: Ces 2 dates correspondent à la période pour laquelle vous souhaitez obtenir une attestation de paiement.</p><p>&nbsp;</p><p>Exemple&nbsp;:</p><p>&nbsp;</p><p>Vous souhaitez une attestation de paiement pour <b>le mois de juillet 2017</b>.</p><p>Il faudra sélectionner comme <b>date de début</b> «&nbsp;Juillet 2017 » &nbsp;et comme <b>date de fin</b> «&nbsp;Juillet 2017 ».</p><p>Vous souhaitez une attestation de paiement pour <b>la période du 1er janvier 2017 au 31 juillet 2017</b>.</p>
<ul>

<li>Il faudra sélectionner comme <b>date de début</b> «&nbsp;Janvier 2017 » et comme<b>&nbsp;date de fin</b> «&nbsp;Juillet 2017 ».</li>

</ul>
<p>&nbsp;</p><a name="eztoc17591585_0_2" id="eztoc17591585_0_2"></a><h3><b>A partir de quand je peux demander une attestation&nbsp;?</b></h3><p>&nbsp;</p><p>Pour les abonnements par prélèvement automatique les attestations de paiement sont disponibles à partir du 2 du mois en cours.</p><p>Exemple&nbsp;: Pour obtenir une attestation pour le mois de juillet 2017, il faudra attendre le 2 juilllet 2017 pour faire votre demande d’attestation de paiement.</p><p>&nbsp;</p><a name="eztoc17591585_0_3" id="eztoc17591585_0_3"></a><h3><b>Délai d’obtention de mon attestation de paiement</b></h3><p>&nbsp;</p><p>Vous obtiendrez par mail votre attestation dans un délai de 72h maximum, vérifiez donc bien que votre adresse mail est correctement renseignée.</p>

					<!-- Module Navitia / lignes_a_proximite bloc bas -->
										<!-- Module Navitia / lignes_a_proximite bloc milieu -->
										<!-- Module Google / google_maps -->
										<!-- Formulaire Mes alertes -->
										<!-- Alertes accessibilité -->
					
					 <!-- Magic node : Page mécanisme colonnes - NE PAS TOUCHER -->
					
									</div>

				<div id="COLONNE-R">

					 

					<!-- Bloc Info trafic accueil -->
					
					
					
					<!-- Bloc Services -->
																							<div class="BLOC-col-D">
                                <!-- Bloc Services -->
                                
                                <div id="encart-itineraire" class="nounderline white">
                                    <a href="/Me-deplacer/Itineraires/Mon-trajet">Calculer mon itinéraire</a>
                                </div>


                                    <div>
                                        <div class="contenu mon-trajet-services">
                                            <ul>
                                                                                                                                                                                                        
                                                                                                                                                                                                            
                                                                                                            <li>


<a href="/Me-deplacer/Toutes-les-lignes"
   target="_self"
   title="Accédez au service Horaires Horaires">
		<span class="picto">
			<img src="/var/tcl/storage/images/media/librairie/encart-service/horaires/360-18-fre-FR/Horaires.png"
                 alt="    Horaires"/>
		</span>
        <span class="text-service">
                Horaires        </span>
</a>


</li>
                                                                                                                                                                                                            
                                                                                                            <li>
            <a href="/Me-deplacer/Plans-du-reseau"
       class="append"
       target="_self"
       title="Accédez au service PDF Plans du réseau">
        <span class="type-file">    PDF</span>
    </a>
    
<a href="http://plan-interactif.tcl.fr/"
   target="_blank"
   title="Accédez au service Plan interactif (nouvelle fenêtre)">
		<span class="picto">
			<img src="/var/tcl/storage/images/media/librairie/encart-service/plan-interactif/374-19-fre-FR/Plan-interactif.png"
                 alt="    Plan interactif"/>
		</span>
        <span class="text-service">
                Plan interactif        </span>
</a>


</li>
                                                                                                                                                                                                            
                                                                                                                                                                                                            
                                                                                                            <li>

    
<a href="https://e-tecely.tcl.fr/accueil.do"
   target="_blank"
   title="Accédez au service Agence en ligne Agence en ligne">
		<span class="picto">
			<img src="/var/tcl/storage/images/media/librairie/encart-service/agence-en-ligne/410-17-fre-FR/Agence-en-ligne.png"
                 alt="    Agence en ligne"/>
		</span>
        <span class="text-service">
                Agence en ligne        </span>
</a>


</li>
                                                                                                                                                                                                            
                                                                                                            <li>


<a href="/Tarifs/Fraude-et-amendes"
   target="_self"
   title="Accédez au service Paiement des amendes Paiement des amendes">
		<span class="picto">
			<img src="/var/tcl/storage/images/media/librairie/encart-service/paiement-des-amendes/381-28-fre-FR/Paiement-des-amendes.png"
                 alt="    Paiement des amendes"/>
		</span>
        <span class="text-service">
                Paiement des amendes        </span>
</a>


</li>
                                                                                                                                                                                                            
                                                                                                                                                                                                </ul>
                                            <div class="CLEAR-both"></div>

                                        </div>
                                        <div class="alertes-accessibilite nounderline">
                                                                        <h3><a href="/Me-deplacer/Alerte-accessibilite">Alertes accessibilité <span>(10)</span></a></h3>
                                                                    </div>
                                    </div>
							</div>
																					<!-- Bloc Actu en bref -->
										<!-- Bloc Guide tarifaire -->
					
					 

					<!-- Bloc Info trafic -->
										<!-- Bloc Horaire Arret -->
										<!-- Bloc Alerte Trafic -->
					

					
					
					<!-- Bloc Parc relais de la ligne -->
										<!-- Bloc Actualité de la ligne -->
										<!-- Bloc Carte tecely  -->
					
					
					 
					
					
					
										<!-- Bloc Mes coordonnées -->
										<!-- Bloc Mes alertes -->
										<!-- Bloc Mes newsletters -->
					
										
					 

					<!-- Bloc Alertes accessibilité -->
										<!-- Bloc Fraude -->
									</div>

				<div class="CLEAR-both"></div>
				

				<!-- FIN CONTENT -->
			</div>
		

			
			
			
			

<div id="FOOTER">
	<div id="BLOC-footer">
				 <div class="BLOC">
		 			 				 <h2>    Contacter TCL</h2>
				 
<p class="NUM-tel">Allô TCL&nbsp;<b>04 26 10 12 12</b></p><p class="NUM-tel-info">Prix d’un appel local</p><p>Du lundi au dimanche : 6h / 22h </p>			 		 			 				 <h2><a class="nous-ecrire" href="/Pied-de-page/Nous-ecrire">    Nous écrire</a></h2>
			 		 		 </div>
				
					<div class="BLOC">
										<h2><a href="/Me-deplacer/Itineraires">Me déplacer</a></h2>

													<ul>
																	<li><a href="/Me-deplacer/Itineraires">Itinéraires</a></li>
																							<li><a href="/Me-deplacer/Toutes-les-lignes">Toutes les lignes</a></li>
																							<li><a href="/Me-deplacer/Infos-trafic">Infos trafic</a></li>
																							<li><a href="/Me-deplacer/Plans-du-reseau">Plans du réseau</a></li>
																							<li><a href="/Me-deplacer/Alerte-accessibilite">Alerte accessibilité</a></li>
																</ul>
														</div>
					<div class="BLOC">
										<h2><a href="/Tarifs/Les-points-de-vente">Tarifs</a></h2>

													<ul>
																	<li><a href="/Tarifs/Les-points-de-vente">Les points de vente</a></li>
																							<li><a href="/Tarifs/S-abonner">S'abonner</a></li>
																							<li><a href="/Tarifs/Tickets">Tickets</a></li>
																							<li><a href='https://e-tecely.tcl.fr' target="_blank" >Agence en ligne</a></li>
																							<li><a href="/Tarifs/Fraude-et-amendes">Fraude et amendes</a></li>
																</ul>
														</div>
					<div class="BLOC">
										<h2><a href="/Decouvrir-TCL/Nouveau-client-TCL">Découvrir TCL</a></h2>

													<ul>
																	<li><a href="/Decouvrir-TCL/Nouveau-client-TCL">Nouveau client TCL</a></li>
																							<li><a href="/Decouvrir-TCL/Le-reseau">Le réseau</a></li>
																							<li><a href="/Decouvrir-TCL/Tous-les-services-TCL">Tous les services TCL</a></li>
																							<li><a href="/Decouvrir-TCL/Accessibilite-du-reseau">Accessibilité du réseau </a></li>
																</ul>
														</div>
				<div class="BLOC">
										<h2>Suivez nous</h2>

				<ul class="list-inline">
									<li>  <a href=https://twitter.com/tcl_sytral> <img src="/var/tcl/storage/images/pied-de-page/suivez-nous/twitter/32041250-1-fre-FR/Twitter.png" alt='Twitter'>  </a></li>
									<li>  <a href=https://www.instagram.com/instatram_tcl/> <img src="/var/tcl/storage/images/pied-de-page/suivez-nous/instagram/32041299-1-fre-FR/Instagram.png" alt='Instagram'>  </a></li>
									<li>  <a href=https://www.facebook.com/TCL.SYTRAL> <img src="/var/tcl/storage/images/pied-de-page/suivez-nous/facebook/32041326-1-fre-FR/Facebook.png" alt='Facebook'>  </a></li>
								</ul>
						<br>
									<h2>Télécharger l'appli</h2>

			<ul class="list-inline">
									<li>  <a href=https://itunes.apple.com/fr/app/tcl/id579379606?mt=8> <img src="/var/tcl/storage/images/pied-de-page/telecharger-l-appli/iphone/32041182-1-fre-FR/Iphone.png" alt='Iphone'>  </a></li>
									<li>  <a href=https://play.google.com/store/apps/details?id=com.micropole.android.tcl_mobile> <img src="/var/tcl/storage/images/pied-de-page/telecharger-l-appli/android/32041239-1-fre-FR/Android.png" alt='Android'>  </a></li>
											</ul>
		</div>

		<a href="#" class="LIEN-haut">HAUT</a>
		<div class="CLEAR-both"></div>
	</div>
	<div id="FOOTER-END">
		
<div  class="logo-footer-end" >
	<p>						


    
        
                                
                        <img              src="/var/tcl/storage/images/media/images/logo-tcl-footer/32042452-1-fre-FR/logo-tcl-footer.png"             width="50"
             height="38"              alt=""
             />
            


    
    			   </p>
</div><div  class="slogan-footer" >
	<p><i>Bougez, Vivez, Aimez !</i></p>
</div><div  class="liens-footer-end" >
	
<ul>

<li>TCL.fr ©2019</li>

<li><a href="http://www.tcl.fr/content/download/1249414/17216300/version/3/file/CHARTE+CLIENTS+TCL.PDF" target="_self">Charte client</a></li>

<li><a href="/Pied-de-page/Infos/Mentions-legales" title="Mentions légales" target="_self">Mentions légales</a></li>

<li><a href="/Pied-de-page/Infos/Plan-du-site" title="Plan du site" target="_self">Plan du Site</a></li>

<li><a href="http://tcl.fr/Pied-de-page/Infos/Participer-au-panel-TCL" target="_self">Le panel TCL</a></li>

</ul>

</div><div  class="certification-footer-end" >
	<p><a href="http://www.afnor.org" title="Accéder au site AFNOR-certification (nouvelle fenêtre)" target="_blank" class="lien-afnor">						


    
        
                                
                        <img              src="/var/tcl/storage/images/media/images/afnor-certification/4155125-1-fre-FR/AFNOR-certification.png"             width="114"
             height="35"              alt="AFNOR-certification"
             />
            


    
    			   </a><a href="http://www.accessiweb.org/index.php/rapport_de_labellisation/items/tcl_2012.html" title="Label AccessiWeb Argent 2012 - Transports en Commun Lyonnais (rapport de labellisation) (nouvelle fenêtre)" target="_blank" class="lien-accessibilite">						


    
        
                                
                        <img              src="/var/tcl/storage/images/media/images/label-accessiweb-argent-2012-transports-en-commun-lyonnais-rapport-de-labellisation/4155128-1-fre-FR/Label-AccessiWeb-Argent-2012-Transports-en-Commun-Lyonnais-rapport-de-labellisation.png"             width="114"
             height="35"              alt="Label AccessiWeb Argent 2012 - Transports en Commun Lyonnais (rapport de labellisation)"
             />
            


    
    			   </a></p>
</div>	</div>
</div><!-- fin FOOTER -->
		</div> <!-- fin CONTAINER -->

																	<!-- eStat -->
		
			<script type="text/javascript">
				<!--
				var _PJS=0;
				//-->
			</script>
			<script type="text/javascript" src="//prof.estat.com/js/276076187740.js"></script>
			<script type="text/javascript">
				<!--
				if(_PJS){
					eStat_id.serial("276076187740");
					eStat_id.master("218018158547");
					eStat_id.gp_pg_mq("index.asp");
					eStat_id.pg_mq("index.asp");
					eStat_tag.post("m");
				}
				//-->
			</script>
			<noscript>
				<img src="http://prof.estat.com/m/web/276076187740?g=218018158547&amp;c=index.asp&amp;p=index.asp&amp;st=0&amp;sjs=0" style="width:1; height:1;" alt="eStat Pro"/>
			</noscript>
			<!-- /eStat -->
		
	</body>
</html>
`