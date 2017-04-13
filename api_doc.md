{% for path, value in sw.Paths.Paths %}
    {{ op.regexReplace(path, "{([^}]+)}", ":$1") }} {% autoescape off %}{{op.prettyjson(value)}} {% endautoescape %}
{% endfor %}